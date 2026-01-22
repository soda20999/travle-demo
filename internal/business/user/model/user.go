// d:/go/project/iam/internal/business/user/model/user.go
package model

import (
	"time"
)
// 修改 Model User 结构体
type User struct {
    ID        int64     `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
    Username  string    `gorm:"column:username;size:255;not null;uniqueIndex" json:"username"`
    Password  string    `gorm:"column:password;size:255;not null" json:"password"`
    Nickname  string    `gorm:"column:nickname;size:64" json:"nickname,omitempty"`
    AvatarURL string    `gorm:"column:avatar_url;size:256" json:"avatar_url,omitempty"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`  
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`  
}
// TableName 指定表名
func (User) TableName() string {
	return "user"
}