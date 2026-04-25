package preference

import (
	"net/http"
	"strconv"


	"github.com/gin-gonic/gin"
)

// CreatePreferenceHandler 创建用户偏好的处理函数
func CreatePreferenceHandler(c *gin.Context) {
	var preference UserTravelPreference
	if err := c.ShouldBindJSON(&preference); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	// 从上下文获取当前用户ID（由JWT中间件设置）
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}
	preference.UserID = uid.(int64)

	if err := CreatePreference(&preference); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

// GetPreferenceHandler 获取用户偏好的处理函数
func GetPreferenceHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}
	userID := uid.(int64)

	preference, err := GetUserPreference(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": preference, "msg": "success"})
}

// GetUserPreferencesHandler 获取用户所有偏好的处理函数
func GetUserPreferencesHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}
	userID := uid.(int64)

	preferences, err := GetUserPreferences(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": preferences, "msg": "success"})
}

// DeletePreferenceHandler 删除用户偏好的处理函数
func DeletePreferenceHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}
	userID := uid.(int64)

	preferredIDStr := c.Param("preferred_id")
	preferredID, err := strconv.ParseInt(preferredIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	if err := DeletePreference(userID, preferredID); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

// GetAllTravelStylesHandler 获取所有旅行风格的处理函数
func GetAllTravelStylesHandler(c *gin.Context) {
	styles, err := GetAllTravelStyles()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": styles, "msg": "success"})
}

// GetTravelStyleByIDHandler 根据ID获取旅行风格的处理函数
func GetTravelStyleByIDHandler(c *gin.Context) {
	styleIDStr := c.Param("style_id")
	styleID, err := strconv.ParseInt(styleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	style, err := GetTravelStyleByID(styleID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": style, "msg": "success"})
}