package model

import (
	"time"

	"gorm.io/datatypes"
)

const (
	ARScanStatusProcessing = 0
	ARScanStatusSuccess    = 1
	ARScanStatusFailed     = 2
)

type ARScan struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID    int64          `gorm:"column:user_id;not null;index" json:"user_id"`
	ImageURL  string         `gorm:"column:image_url;size:500;not null" json:"image_url"`
	Status    int            `gorm:"column:status;not null;default:0;index" json:"status"`
	Metadata  datatypes.JSON `gorm:"column:metadata;type:jsonb" json:"metadata"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Results []ARScanResult `gorm:"foreignKey:ScanID;references:ID" json:"results,omitempty"`
}

type ARScanResult struct {
	ID         int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ScanID     int64          `gorm:"column:scan_id;not null;index" json:"scan_id"`
	ModelName  string         `gorm:"column:model_name;size:100;not null" json:"model_name"`
	ObjectName string         `gorm:"column:object_name;size:255;not null" json:"object_name"`
	BriefInfo  string         `gorm:"column:brief_info;type:text" json:"brief_info"`
	FullData   datatypes.JSON `gorm:"column:full_data;type:jsonb" json:"full_data"`
	Confidence float64        `gorm:"column:confidence;type:numeric(4,3)" json:"confidence"`
	CreatedAt  time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Scan *ARScan `gorm:"foreignKey:ScanID;references:ID" json:"scan,omitempty"`
}

func (ARScan) TableName() string {
	return "ar_scans"
}

func (ARScanResult) TableName() string {
	return "ar_scan_results"
}
