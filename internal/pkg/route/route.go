package route

import (
	ar_api "iam/internal/business/ar/api"
	discover_api "iam/internal/business/discover/api"
	footprint_api "iam/internal/business/footprint/api"
	preference_api "iam/internal/business/preference/api"
	rec_api "iam/internal/business/recognize/api"
	user_api "iam/internal/business/user/api"
	"iam/internal/pkg/config/logger"
	"iam/internal/pkg/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 1. 跨域中间件 (极简版)
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 公开路由 (Public) ---
	r.POST("/signup", user_api.SignHandler)
	r.POST("/login", user_api.LoginHandler)
	r.POST("/refresh_token", user_api.RefreshTokenHandler)
	r.GET("/travel-styles", preference_api.GetAllTravelStylesHandler)
	r.GET("/travel-styles/:style_id", preference_api.GetTravelStyleByIDHandler)

	// 发现模块：地理信息 (通常不需要登录也能看)
	discover := r.Group("/")
	{
		discover.GET("/provinces", discover_api.GetProvincesHandler)
		discover.GET("/provinces/:province_id", discover_api.GetProvinceByIDHandler)
		discover.GET("/provinces/:province_id/cities", discover_api.GetCitiesByProvinceHandler)
		discover.GET("/cities/:city_id", discover_api.GetCityByIDHandler)
		discover.GET("/attractions/:attraction_id", discover_api.GetAttractionByIDHandler)
	}

	// --- 认证路由 (Protected) ---
	// 使用 Group 管理需要 JWT 的接口，避免重复写 middlewares.JWTAuthMiddleware()
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		// 用户个人信息
		user := auth.Group("/user")
		{
			user.GET("/info", user_api.GetUserInfoHandler)
			user.POST("/update-nickname", user_api.UpdateNicknameHandler)
			user.POST("/update-avatar", user_api.UpdateAvatarHandler)
		}

		// 偏好设置
		pref := auth.Group("/preferences")
		{
			pref.POST("", preference_api.CreatePreferenceHandler)
			pref.GET("", preference_api.GetUserPreferencesHandler)
			pref.GET("/:user_id", preference_api.GetPreferenceHandler)
			pref.DELETE("/:preferred_id", preference_api.DeletePreferenceHandler)
		}

		// 足迹相关
		foot := auth.Group("/footprints")
		{
			foot.GET("", footprint_api.GetFootprintsHandler)
			foot.POST("", footprint_api.CreateFootprintHandler)
			foot.DELETE("/:id", footprint_api.DeleteFootprintHandler)
		}

		// AR相关
		ar := auth.Group("/ar-scans")
		{
			ar.GET("", ar_api.GetARScansHandler)
			ar.GET("/:id", ar_api.GetARScanByIDHandler)
			ar.POST("", ar_api.CreateARScanHandler)
		}

		auth.POST("/recognize", rec_api.RecognizeHandler)

		gallery := auth.Group("/gallery")
		{
			gallery.POST("/images", rec_api.AddGalleryImageHandler)
			gallery.DELETE("/images/:id", rec_api.DeleteGalleryImageHandler)
			gallery.GET("/attractions/:id/images", rec_api.GetGalleryImagesHandler)
			gallery.POST("/rebuild-index", rec_api.RebuildIndexHandler)
		}

		auth.GET("/home", middlewares.HomeHandler)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
