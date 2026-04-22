package footprint_service

import (
	"time"

	foot_model "iam/internal/business/footprint/model"
	"iam/internal/pkg/config/postsql"
)

// FootprintWithAttraction 用于返回足迹及关联的景点信息（名称、图片）
// 注意：这个结构仅用于内部组装响应数据，不是 vo，也不导出到 API 层
type footprintWithAttraction struct {
	foot_model.Footprint
	Name  string
	Image string
}

// GetFootprintsByUserID 获取用户的足迹列表，并填充景点名称和图片
func GetFootprintsByUserID(userID int64) ([]footprintWithAttraction, error) {
	var footprints []foot_model.Footprint
	err := postgresql.DB.Preload("Attraction").Where("user_id = ?", userID).Order("created_at DESC").Find(&footprints).Error
	if err != nil {
		return nil, err
	}

	result := make([]footprintWithAttraction, 0, len(footprints))
	for _, fp := range footprints {
		item := footprintWithAttraction{
			Footprint: fp,
		}
		if fp.Attraction != nil {
			item.Name = fp.Attraction.Name
			item.Image = fp.Attraction.Image
		}
		result = append(result, item)
	}
	return result, nil
}

// CreateFootprint 创建足迹
func CreateFootprint(userID, attractionID int64, date string) error {
	footprint := &foot_model.Footprint{
		UserID:       userID,
		AttractionID: attractionID,
		Date:         date,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return postgresql.DB.Create(footprint).Error
}

// DeleteFootprint 删除足迹
func DeleteFootprint(id int64) error {
	return postgresql.DB.Where("id = ?", id).Delete(&foot_model.Footprint{}).Error
}