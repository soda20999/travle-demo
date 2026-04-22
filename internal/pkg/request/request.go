package request

import (
	"errors"

	"iam/internal/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var ErrUserNotLogin = errors.New("用户未登录")

// GetUserID 从Gin的Context中获取当前请求的用户ID
func GetUserID(c *gin.Context) (int64, error) {
	uid, ok := c.Get(middlewares.ContextUserID)
	if !ok {
		return 0, ErrUserNotLogin
	}
	userID, ok := uid.(int64)
	if !ok {
		return 0, ErrUserNotLogin
	}
	return userID, nil
}
