package vo

import "github.com/go-playground/validator/v10"

// ParamCreatePreference 创建偏好参数
type ParamCreatePreference struct {
    UserID      int64 `json:"user_id" binding:"required"`
    PreferredID int64 `json:"preferred_id" binding:"required"`
}

// ParamGetPreference 查询偏好参数
type ParamGetPreference struct {
    UserID int64 `uri:"user_id" binding:"required"`
}

// ResponsePreference 偏好响应
type ResponsePreference struct {
    UserID    int64 `json:"user_id"`
    StyleName string `json:"style_name"`
    StyleID   int64 `json:"style_id"`
}

// ResponsePreferenceList 偏好列表响应
type ResponsePreferenceList struct {
    Preferences []ResponsePreference `json:"preferences"`
}

// Validate 验证函数
func (p *ParamCreatePreference) Validate() error {
    validate := validator.New()
    return validate.Struct(p)
}