// internal/business/preference/repo/preference.go
package repo

import (
	"errors"
	"iam/internal/business/preference/model"
	"iam/internal/pkg/config/gorm"
)

var (
	ErrorPreferenceNotFound = errors.New("偏好设置未找到")
	ErrorStyleNotFound      = errors.New("旅行风格未找到")
)

// CreatePreference 创建用户偏好
func CreatePreference(preference *model.UserTravelPreference) error {
	result := gorm.Db.Create(preference)
	return result.Error
}

// GetPreferenceByUserID 根据用户ID获取偏好
func GetPreferenceByUserID(userID int64) (*model.UserTravelPreference, error) {
	var preference model.UserTravelPreference
	result := gorm.Db.Preload("TravelStyle").Where("user_id = ?", userID).First(&preference)
	
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrorPreferenceNotFound
		}
		return nil, result.Error
	}
	
	return &preference, nil
}

// GetUserPreferences 获取用户的所有偏好
func GetUserPreferences(userID int64) ([]model.UserTravelPreference, error) {
	var preferences []model.UserTravelPreference
	result := gorm.Db.Preload("TravelStyle").
		Where("user_id = ?", userID).
		Find(&preferences)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return preferences, nil
}

// DeletePreference 删除用户偏好
func DeletePreference(userID int64, preferredID int64) error {
	result := gorm.Db.Where("user_id = ? AND preferred_id = ?", userID, preferredID).Delete(&model.UserTravelPreference{})
	return result.Error
}

// GetAllTravelStyles 获取所有旅行风格
func GetAllTravelStyles() ([]model.TravelStyle, error) {
	var styles []model.TravelStyle
	result := gorm.Db.Find(&styles)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return styles, nil
}

// GetTravelStyleByID 根据ID获取旅行风格
func GetTravelStyleByID(styleID int64) (*model.TravelStyle, error) {
	var style model.TravelStyle
	result := gorm.Db.Where("id = ?", styleID).First(&style)
	
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrorStyleNotFound
		}
		return nil, result.Error
	}
	
	return &style, nil
}