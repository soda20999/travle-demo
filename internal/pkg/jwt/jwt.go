package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("WASDMu123com")
var ErrorTokenExpired = errors.New("token不存在")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(username string, userID int64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		UserID:   userID,
		Username: username, // 自定义字段
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    "iam",                                                   //app名称
		},
	}

	//access token 加密并获得完整的编码后的字符串串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256,
		claims).SignedString(CustomSecret)
	// refresh token 不不需要存任何⾃自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)), // 过期时间
		Issuer:    "iam",                                                // 签发人
	}).SignedString(CustomSecret)
	// 使⽤用指定的secret签名并获得完整的编码后的字符
	return aToken, rToken, err

}

/*
// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
*/

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *CustomClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(CustomClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = ErrorTokenExpired
	}
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	// 从旧access token中解析出claims数据
	var claims CustomClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	// 判断是否为token过期错误
	if errors.Is(err, jwt.ErrTokenExpired) {
		// 过期则生成新token
		return GenToken(claims.Username, claims.UserID)
	}

	return
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return CustomSecret, nil
}
