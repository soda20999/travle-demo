package ar_service

import (
	"errors"

	ar_model "iam/internal/business/ar/model"
	postgresql "iam/internal/pkg/config/postsql"

	"gorm.io/gorm"
)

func CreateARScan(scan *ar_model.ARScan) error {
	return postgresql.DB.Create(scan).Error
}

func GetARScansByUserID(userID int64) ([]ar_model.ARScan, error) {
	var scans []ar_model.ARScan
	err := postgresql.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&scans).Error
	return scans, err
}

func UpdateARScanStatus(scanID int64, status int) error {
	return postgresql.DB.Model(&ar_model.ARScan{}).
		Where("id = ?", scanID).
		Update("status", status).Error
}

func GetARScanByID(scanID int64) (*ar_model.ARScan, error) {
	var scan ar_model.ARScan
	err := postgresql.DB.Preload("Results").Where("id = ?", scanID).First(&scan).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &scan, nil
}

func CreateARScanResult(result *ar_model.ARScanResult) error {
	return postgresql.DB.Create(result).Error
}

func GetARScanResultsByScanID(scanID int64) ([]ar_model.ARScanResult, error) {
	var results []ar_model.ARScanResult
	err := postgresql.DB.Where("scan_id = ?", scanID).Order("created_at DESC").Find(&results).Error
	return results, err
}

func GetLatestARScanResult(scanID int64) (*ar_model.ARScanResult, error) {
	var result ar_model.ARScanResult
	err := postgresql.DB.Where("scan_id = ?", scanID).Order("created_at DESC").First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
