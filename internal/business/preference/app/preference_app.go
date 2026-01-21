package app

import (
    "errors"
    "iam/internal/business/preference/domin/dto"
    "iam/internal/business/preference/model"
    "iam/internal/business/preference/repo"
    "iam/pkg/snowflake"
)

// CreatePreference 创建用户偏好
func CreatePreference(param *dto.Preference) error {
    // 验证旅行风格是否存在
    _, err := repo.GetTravelStyleByID(param.PreferredID)
    if err != nil {
        return err
    }
    
    // 检查用户是否已有相同的偏好
    existingPref, _ := repo.GetPreferenceByUserID(param.UserID)
    if existingPref != nil && existingPref.PreferredID == param.PreferredID {
        return errors.New("用户已存在相同的偏好设置")
    }
    
    preference := &model.UserTravelPreference{
        ID:          snowflake.GenID(),
        UserID:      param.UserID,
        PreferredID: param.PreferredID,
    }
    
    return repo.CreatePreference(preference)
}

// GetUserPreference 获取用户偏好
func GetUserPreference(userID int64) (*dto.PreferenceWithStyle, error) {
    pref, err := repo.GetPreferenceByUserID(userID)
    if err != nil {
        return nil, err
    }
    
    if pref.TravelStyle == nil {
        return &dto.PreferenceWithStyle{
            UserID:    pref.UserID,
            StyleName: "",
            StyleID:   pref.PreferredID,
        }, nil
    }
    
    return &dto.PreferenceWithStyle{
        UserID:    pref.UserID,
        StyleName: pref.TravelStyle.StyleName,
        StyleID:   pref.PreferredID,
    }, nil
}

// GetUserPreferences 获取用户的所有偏好
func GetUserPreferences(userID int64) ([]dto.PreferenceWithStyle, error) {
    prefs, err := repo.GetUserPreferences(userID)
    if err != nil {
        return nil, err
    }
    
    var result []dto.PreferenceWithStyle
    for _, pref := range prefs {
        styleName := ""
        if pref.TravelStyle != nil {
            styleName = pref.TravelStyle.StyleName
        }
        
        result = append(result, dto.PreferenceWithStyle{
            UserID:    pref.UserID,
            StyleName: styleName,
            StyleID:   pref.PreferredID,
        })
    }
    
    return result, nil
}

// DeletePreference 删除用户偏好
func DeletePreference(userID int64, preferredID int64) error {
    return repo.DeletePreference(userID, preferredID)
}

// GetAllTravelStyles 获取所有旅行风格
func GetAllTravelStyles() ([]model.TravelStyle, error) {
    return repo.GetAllTravelStyles()
}