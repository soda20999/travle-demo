package route

import (
	"net/http"

	preference_api "iam/internal/business/preference/api"
	user_api "iam/internal/business/user/api"
	"iam/internal/pkg/config/logger"
	"iam/internal/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//1.注册业务路由
	r.POST("/signup", user_api.SignHandler)
	r.POST("/login", user_api.LoginHandler)

	//2.用户偏好路由
	r.POST("/preferences", middlewares.JWTAuthMiddleware(), preference_api.CreatePreferenceHandler)
	r.GET("/preferences", middlewares.JWTAuthMiddleware(), preference_api.GetUserPreferencesHandler)
	r.GET("/preferences/:user_id", middlewares.JWTAuthMiddleware(), preference_api.GetPreferenceHandler)
	r.DELETE("/preferences/:preferred_id", middlewares.JWTAuthMiddleware(), preference_api.DeletePreferenceHandler)
	r.GET("/travel-styles", preference_api.GetAllTravelStylesHandler) // 这个可以公开访问



	//使用JWT中间件保护需要认证的路由
	r.GET("/home", middlewares.JWTAuthMiddleware(), middlewares.HomeHandler)
	r.POST("/refresh_token", user_api.RefreshTokenHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
