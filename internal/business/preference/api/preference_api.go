package api

import (
    "strconv"
    
    "iam/internal/business/preference/app"
    "iam/internal/business/preference/domin/dto"
    "iam/internal/business/preference/domin/vo"
    "iam/internal/pkg/request"
    "iam/internal/pkg/response"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// CreatePreferenceHandler 创建用户偏好的处理函数
func CreatePreferenceHandler(c *gin.Context) {
    p := new(vo.ParamCreatePreference)
    if err := c.ShouldBindJSON(p); err != nil {
        zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
        zap.L().Error("create preference with invalid param")
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }
    
    // 从请求上下文获取当前用户的ID
    userID, err := request.GetUserID(c)
    if err != nil {
        response.ResponseError(response.CodeNeedLogin, c)
        return
    }
    p.UserID = userID
    
    // 构造DTO
    preference := &dto.Preference{
        UserID:      p.UserID,
        PreferredID: p.PreferredID,
    }
    
    // 调用应用层创建偏好
    if err := app.CreatePreference(preference); err != nil {
        zap.L().Error("app.CreatePreference failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }
    
    response.ResponseSuccess(c, nil)
}

// GetPreferenceHandler 获取用户偏好的处理函数
func GetPreferenceHandler(c *gin.Context) {
    userIDStr := c.Param("user_id")
    userID, err := strconv.ParseInt(userIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }
    
    // 权限检查：只能查看自己的偏好
    currentUserID, err := request.GetUserID(c)
    if err != nil || currentUserID != userID {
        response.ResponseError(response.CodeNeedLogin, c)
        return
    }
    
    preference, err := app.GetUserPreference(userID)
    if err != nil {
        zap.L().Error("app.GetUserPreference failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }
    
    response.ResponseSuccess(c, preference)
}

// GetUserPreferencesHandler 获取用户所有偏好的处理函数
func GetUserPreferencesHandler(c *gin.Context) {
    userID, err := request.GetUserID(c)
    if err != nil {
        response.ResponseError(response.CodeNeedLogin, c)
        return
    }
    
    preferences, err := app.GetUserPreferences(userID)
    if err != nil {
        zap.L().Error("app.GetUserPreferences failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }
    
    response.ResponseSuccess(c, vo.ResponsePreferenceList{
        Preferences: convertToResponse(preferences),
    })
}

// DeletePreferenceHandler 删除用户偏好的处理函数
func DeletePreferenceHandler(c *gin.Context) {
    userID, err := request.GetUserID(c)
    if err != nil {
        response.ResponseError(response.CodeNeedLogin, c)
        return
    }
    
    preferredIDStr := c.Param("preferred_id")
    preferredID, err := strconv.ParseInt(preferredIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }
    
    if err := app.DeletePreference(userID, preferredID); err != nil {
        zap.L().Error("app.DeletePreference failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }
    
    response.ResponseSuccess(c, nil)
}

// GetAllTravelStylesHandler 获取所有旅行风格的处理函数
func GetAllTravelStylesHandler(c *gin.Context) {
    styles, err := app.GetAllTravelStyles()
    if err != nil {
        zap.L().Error("app.GetAllTravelStyles failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }
    
    response.ResponseSuccess(c, styles)
}

// 辅助函数：转换为响应格式
func convertToResponse(prefs []dto.PreferenceWithStyle) []vo.ResponsePreference {
    result := make([]vo.ResponsePreference, len(prefs))
    for i, pref := range prefs {
        result[i] = vo.ResponsePreference{
            UserID:    pref.UserID,
            StyleName: pref.StyleName,
            StyleID:   pref.StyleID,
        }
    }
    return result
}