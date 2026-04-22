package api

import (
    "github.com/gin-gonic/gin"
    "iam/internal/business/user/model"
    "iam/internal/business/user/service"
    "net/http"
	"iam/internal/pkg/jwt"
)
// 1. 注册 (POST /signup)
func SignHandler(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}
	if err := user_service.SignUp(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10004, "msg": "用户已存在或服务器繁忙"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

func LoginHandler(c *gin.Context) {
    var u model.User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
        return
    }

    user, err := user_service.Login(u.Username, u.Password)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 10006, "msg": "用户名或密码错误"})
        return
    }

    // --- 关键修改：对接你自己的 jwt.GenToken ---
    aToken, rToken, err := jwt.GenToken(user.Username, user.ID)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "data": gin.H{
            "access_token":  aToken,
            "refresh_token": rToken,
        },
        "msg": "success",
    })
}
// 4. 更新昵称 (POST /user/update-nickname)
func UpdateNicknameHandler(c *gin.Context) {
	var req struct {
		UserID   int64  `json:"user_id" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}
	if err := user_service.UpdateUser(req.UserID, map[string]interface{}{"nickname": req.Nickname}); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

// 5. 更新头像 (POST /user/update-avatar)
func UpdateAvatarHandler(c *gin.Context) {
	var req struct {
		UserID    int64  `json:"user_id" binding:"required"`
		AvatarURL string `json:"avatar_url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}
	if err := user_service.UpdateUser(req.UserID, map[string]interface{}{"avatar_url": req.AvatarURL}); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

// 6. 获取用户信息 (GET /user/info)
func GetUserInfoHandler(c *gin.Context) {
	// 从 JWT 中间件提取解析出的 userID
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}

	u, err := user_service.GetUserInfo(uid.(int64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}
	
	// Model 里的 Password 字段因为有 json:"-" 标签，这里序列化后不会发给前端
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": u, "msg": "success"})
}