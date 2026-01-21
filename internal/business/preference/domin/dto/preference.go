package dto

// Preference 用户偏好数据传输对象
type Preference struct {
    UserID      int64 `json:"user_id"`
    PreferredID int64 `json:"preferred_id"`
}

// PreferenceWithStyle 包含旅行风格信息的偏好
type PreferenceWithStyle struct {
    UserID    int64 `json:"user_id"`
    StyleName string `json:"style_name"`
    StyleID   int64 `json:"style_id"`
}