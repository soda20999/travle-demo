// internal/business/recognize/model/recognize.go
package model

import "time"

type AttractionImage struct {
	ID            int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	AttractionID  int64     `gorm:"column:attraction_id;not null;index" json:"attraction_id"`
	ImagePath     string    `gorm:"column:image_path;size:500;not null" json:"image_path"`
	FeatureVector []byte    `gorm:"column:feature_vector;type:bytea" json:"-"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (AttractionImage) TableName() string { return "attraction_images" }
