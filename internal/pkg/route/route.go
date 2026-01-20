package route

import (
	"net/http"

	"iam/internal/business/user/api"
	"iam/internal/pkg/config/logger"
	"iam/internal/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//1.注册业务路由
	r.POST("/signup", api.SignHandler)
	r.POST("/login", api.LoginHandler)

	//使用JWT中间件保护需要认证的路由
	r.GET("/home", middlewares.JWTAuthMiddleware(), middlewares.HomeHandler)
	r.POST("/refresh_token", api.RefreshTokenHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
