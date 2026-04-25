package preference

import (
	"errors"

	"iam/internal/pkg/config/postsql"

	"gorm.io/gorm"
)

// CreatePreference 创建用户偏好
func CreatePreference(preference *UserTravelPreference) error {
	return postgresql.DB.Create(preference).Error
}

// GetUserPreference 根据用户ID获取偏好（单条）
func GetUserPreference(userID int64) (*UserTravelPreference, error) {
	var pref UserTravelPreference
	err := postgresql.DB.Where("user_id = ?", userID).First(&pref).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 未找到返回 nil，不报错
		}
		return nil, err
	}
	return &pref, nil
}

// GetUserPreferences 获取用户的所有偏好（列表）
func GetUserPreferences(userID int64) ([]UserTravelPreference, error) {
	var prefs []UserTravelPreference
	err := postgresql.DB.Where("user_id = ?", userID).Find(&prefs).Error
	return prefs, err
}

// DeletePreference 删除用户偏好
func DeletePreference(userID int64, preferredID int64) error {
	return postgresql.DB.Where("user_id = ? AND preferred_id = ?", userID, preferredID).Delete(&UserTravelPreference{}).Error
}

// GetAllTravelStyles 获取所有旅行风格
func GetAllTravelStyles() ([]TravelStyle, error) {
	var styles []TravelStyle
	err := postgresql.DB.Find(&styles).Error
	return styles, err
}

// GetTravelStyleByID 根据ID获取旅行风格
func GetTravelStyleByID(styleID int64) (*TravelStyle, error) {
	var style TravelStyle
	err := postgresql.DB.Where("id = ?", styleID).First(&style).Error
	if err != nil {
		return nil, err
	}
	return &style, nil
}