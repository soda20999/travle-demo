package ar

import (
	"errors"

	"iam/internal/pkg/config/postsql"

	"gorm.io/gorm"
)

func CreateARScan(scan *ARScan) error {
	return postgresql.DB.Create(scan).Error
}

func GetARScansByUserID(userID int64) ([]ARScan, error) {
	var scans []ARScan
	err := postgresql.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&scans).Error
	return scans, err
}

func UpdateARScanStatus(scanID int64, status int) error {
	return postgresql.DB.Model(&ARScan{}).
		Where("id = ?", scanID).
		Update("status", status).Error
}

func GetARScanByID(scanID int64) (*ARScan, error) {
	var scan ARScan
	err := postgresql.DB.Preload("Results").Where("id = ?", scanID).First(&scan).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &scan, nil
}

func CreateARScanResult(result *ARScanResult) error {
	return postgresql.DB.Create(result).Error
}

func GetARScanResultsByScanID(scanID int64) ([]ARScanResult, error) {
	var results []ARScanResult
	err := postgresql.DB.Where("scan_id = ?", scanID).Order("created_at DESC").Find(&results).Error
	return results, err
}

func GetLatestARScanResult(scanID int64) (*ARScanResult, error) {
	var result ARScanResult
	err := postgresql.DB.Where("scan_id = ?", scanID).Order("created_at DESC").First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
