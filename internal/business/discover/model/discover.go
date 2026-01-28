package model

import (
    "time"
)

// Province 省份模型
type Province struct {
    ID        int64     `gorm:"primaryKey;autoIncrement;column:id"`
    Name      string    `gorm:"column:name;size:50;not null;uniqueIndex"`
    Code      string    `gorm:"column:code;size:10;uniqueIndex"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`

    Cities []City `gorm:"foreignKey:ProvinceID"`
}

// City 城市模型
type City struct {
    ID         int64     `gorm:"primaryKey;autoIncrement;column:id"`
    Name       string    `gorm:"column:name;size:50;not null"`
    ProvinceID int64     `gorm:"column:province_id;not null"`
    Code       string    `gorm:"column:code;size:10;uniqueIndex"`
    CreatedAt  time.Time `gorm:"column:created_at"`
    UpdatedAt  time.Time `gorm:"column:updated_at"`

    Province     *Province       `gorm:"foreignKey:ProvinceID"`
    Attractions  []Attraction    `gorm:"foreignKey:CityID"`
}

// Attraction 景点模型
type Attraction struct {
    ID        int64     `gorm:"primaryKey;autoIncrement;column:id"`
    Name      string    `gorm:"column:name;size:100;not null"`
    CityID    int64     `gorm:"column:city_id;not null"`
    Code      string    `gorm:"column:code;size:10;uniqueIndex"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`

    City *City `gorm:"foreignKey:CityID"`
}

// TableName 指定表名
func (Province) TableName() string {
    return "travel_provinces"
}

func (City) TableName() string {
    return "travel_cities"
}

func (Attraction) TableName() string {
    return "travel_attractions"
}