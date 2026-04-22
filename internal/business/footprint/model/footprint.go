package model

import (
    "iam/internal/business/discover/model"
    "time"
)

type Footprint struct {
    ID           int64             `gorm:"primaryKey;autoIncrement;column:id"`
    UserID       int64             `gorm:"column:user_id;not null;index"`
    AttractionID int64             `gorm:"column:attraction_id;not null"`
    Date         string            `gorm:"column:date;size:20"`
    CreatedAt    time.Time         `gorm:"column:created_at"`
    UpdatedAt    time.Time         `gorm:"column:updated_at"`

    Attraction *model.Attraction `gorm:"foreignKey:AttractionID"`
}

func (Footprint) TableName() string {
    return "user_footprints"
}
