// d:/go/project/iam/internal/business/user/repo/user.go
package repo

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"iam/internal/business/user/domin/dto"
	"iam/internal/business/user/model"
	"iam/internal/pkg/config/gorm"
)

const secretKey = "mumu.123.com"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// QueryUserByUserName 检查用户是否存在
func QueryUserByUserName(username string) error {
	var user model.User
	result := gorm.Db.Where("username = ?", username).First(&user)
	
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil // 用户不存在，不是错误
		}
		return result.Error // 数据库错误
	}
	
	return ErrorUserExist // 用户已存在
}

// InsertUser 插入用户信息
func InsertUser(user dto.User) error {
	modelUser := user.ToModel()
	modelUser.Password = encryptPassword(modelUser.Password)
	
	result := gorm.Db.Create(modelUser)
	return result.Error
}

// Login 根据用户名和密码登录
func Login(user *dto.User) (err error) {
	var modelUser model.User
	result := gorm.Db.Where("username = ?", user.Username).First(&modelUser)
	
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ErrorUserNotExist
		}
		return result.Error // 数据库错误
	}

	// 验证密码是否错误
	password := encryptPassword(user.Password)
	if password != modelUser.Password {
		return ErrorInvalidPassword
	}
	
	// 更新 DTO 信息
	user.UserID = modelUser.ID
	user.Password = modelUser.Password
	
	return nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password + secretKey))
	return hex.EncodeToString(h.Sum([]byte(password)))
}