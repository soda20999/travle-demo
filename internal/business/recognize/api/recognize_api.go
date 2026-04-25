package api

import (
	"net/http"
	"strconv"

	rec_model "iam/internal/business/recognize/model"
	rec_service "iam/internal/business/recognize/service"

	"github.com/gin-gonic/gin"
)

func RecognizeHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "缺少图片参数"})
		return
	}

	topK := 5
	if v := c.PostForm("top_k"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed > 0 && parsed <= 20 {
			topK = parsed
		}
	}

	results, err := rec_service.Recognize(file, topK)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "识别服务异常"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": results,
		"msg":  "success",
	})
}

func AddGalleryImageHandler(c *gin.Context) {
	var req struct {
		AttractionID int64  `json:"attraction_id" binding:"required"`
		ImageURL     string `json:"image_url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	img := rec_model.AttractionImage{
		AttractionID: req.AttractionID,
		ImageURL:     req.ImageURL,
	}
	if err := rec_service.CreateAttractionImage(&img); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "添加失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":            img.ID,
			"attraction_id": img.AttractionID,
			"image_url":     img.ImageURL,
			"created_at":    img.CreatedAt,
		},
		"msg": "success",
	})
}

func DeleteGalleryImageHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	if err := rec_service.DeleteAttractionImage(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

func GetGalleryImagesHandler(c *gin.Context) {
	idStr := c.Param("id")
	attractionID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	images, err := rec_service.GetImagesByAttractionID(attractionID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": images,
		"msg":  "success",
	})
}

func RebuildIndexHandler(c *gin.Context) {
	if err := rec_service.RebuildIndex(); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "重建索引失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}
