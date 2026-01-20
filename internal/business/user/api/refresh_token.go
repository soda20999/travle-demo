package api

import (
	"iam/internal/pkg/jwt"
	"iam/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type RefreshTokenRequest struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func RefreshTokenHandler(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseErrorWithMsg("参数错误", c)
		return
	}

	newAToken, newRToken, err := jwt.RefreshToken(req.AccessToken, req.RefreshToken)
	if err != nil {
		response.ResponseErrorWithMsg("refresh token无效或已过期，请重新登录", c)
		return
	}

	response.ResponseSuccess(gin.H{
		"access_token":  newAToken,
		"refresh_token": newRToken,
	}, c)
}
