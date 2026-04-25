package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	ar_model "iam/internal/business/ar/model"
	ar_service "iam/internal/business/ar/service"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func CreateARScanHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}

	var req struct {
		ImageURL string          `json:"image_url" binding:"required"`
		Metadata json.RawMessage `json:"metadata"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	scan := ar_model.ARScan{
		UserID:   uid.(int64),
		ImageURL: req.ImageURL,
		Status:   ar_model.ARScanStatusProcessing,
	}
	if len(req.Metadata) > 0 {
		scan.Metadata = datatypes.JSON(req.Metadata)
	}

	if err := ar_service.CreateARScan(&scan); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":         scan.ID,
			"status":     scan.Status,
			"image_url":  scan.ImageURL,
			"metadata":   scan.Metadata,
			"created_at": scan.CreatedAt,
		},
		"msg": "success",
	})
}

func GetARScansHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}

	scans, err := ar_service.GetARScansByUserID(uid.(int64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": scans, "msg": "success"})
}

func GetARScanByIDHandler(c *gin.Context) {
	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "需要登录"})
		return
	}

	scanID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	scan, err := ar_service.GetARScanByID(scanID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}
	if scan == nil || scan.UserID != uid.(int64) {
		c.JSON(http.StatusOK, gin.H{"code": 10005, "msg": "记录不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": scan, "msg": "success"})
}
