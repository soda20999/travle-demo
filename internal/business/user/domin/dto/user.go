// d:/go/project/iam/internal/business/user/domin/dto/user.go
package dto

import (
	"iam/internal/business/user/model"
)

type User struct {
    UserID   int64  `json:"user_id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Nickname string `json:"nickname,omitempty"`
    AvatarURL string `json:"avatar_url,omitempty"`
}

// 转换函数：DTO -> Model
func (u *User) ToModel() *model.User {
	return &model.User{
		ID:       u.UserID,
		Username: u.Username,
		Password: u.Password,
	}
}

// 转换函数：Model -> DTO
func FromModel(m *model.User) *User {
    return &User{
        UserID:   m.ID,
        Username: m.Username,
        Password: m.Password,
        Nickname: m.Nickname,
        AvatarURL: m.AvatarURL,
    }
}