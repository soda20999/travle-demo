package api

import (
	"net/http"
	"iam/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)
// 3. 刷新 Token (POST /refresh_token)
func RefreshTokenHandler(c *gin.Context) {
    // 使用匿名结构体，极致去 DTO
    var req struct {
        AccessToken  string `json:"access_token" binding:"required"`
        RefreshToken string `json:"refresh_token" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
        return
    }

    // 调用你 pkg/jwt 下的刷新逻辑
    newAToken, newRToken, err := jwt.RefreshToken(req.AccessToken, req.RefreshToken)
    if err != nil {
        // 10003 通常代表认证失效，需重新登录
        c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "身份验证已过期，请重新登录"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "data": gin.H{
            "access_token":  newAToken,
            "refresh_token": newRToken,
        },
        "msg": "success",
    })
}