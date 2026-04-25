package footprint

import (
	"net/http"
	"strconv"


	"github.com/gin-gonic/gin"
)

// GetFootprintsHandler 获取用户足迹列表
func GetFootprintsHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}
	userID := uid.(int64)

	footprints, err := GetFootprintsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	// 手动构造返回列表（不使用 vo）
	list := make([]gin.H, 0, len(footprints))
	for _, fp := range footprints {
		list = append(list, gin.H{
			"id":    fp.ID,
			"date":  fp.Date,
			"name":  fp.Name,
			"image": fp.Image,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"footprints": list},
		"msg":  "success",
	})
}

// CreateFootprintHandler 创建用户足迹
func CreateFootprintHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}
	userID := uid.(int64)

	var params struct {
		AttractionID int64  `json:"attraction_id" binding:"required"`
		Date         string `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	if err := CreateFootprint(userID, params.AttractionID, params.Date); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

// DeleteFootprintHandler 删除用户足迹
func DeleteFootprintHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	if err := DeleteFootprint(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}