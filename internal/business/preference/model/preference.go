package model

import (
    "time"
)

// UserTravelPreference 用户旅行偏好模型
type UserTravelPreference struct {
    ID          int64      `gorm:"primaryKey;autoIncrement;column:id"`
    UserID      int64      `gorm:"column:user_id;unique;not null"`
    PreferredID int64      `gorm:"column:preferred_id"`
    CreatedAt   time.Time   `gorm:"column:created_at"`
    UpdatedAt   time.Time   `gorm:"column:updated_at"`
    
    // 关联的旅行风格偏好
    TravelStyle *TravelStyle `gorm:"foreignKey:PreferredID;references:ID"`
}

// TravelStyle 旅行风格偏好模型
type TravelStyle struct {
    ID          int64    `gorm:"primaryKey;autoIncrement;column:id"`
    StyleName   string    `gorm:"column:style_name;size:50;not null"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName 指定表名
func (UserTravelPreference) TableName() string {
    return "user_travel_preferences"
}

func (TravelStyle) TableName() string {
    return "preference_travel_styles"
}