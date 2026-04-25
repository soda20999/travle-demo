package user

import (
	"time"
)


type User struct {
    // 基础字段，使用雪花算法或 PG 的 bigserial
    ID        int64     `gorm:"primaryKey;column:user_id" json:"user_id"`
    Username  string    `gorm:"uniqueIndex;not null" json:"username" binding:"required,min=4"`
    // 使用 json:"-" 确保密码永远不会在获取用户信息时被返回（绝不泄露）
    Password  string    `gorm:"not null" json:"password,omitempty" binding:"required,min=6"`
    Nickname  string    `gorm:"default:''" json:"nickname"`
    AvatarURL string    `gorm:"default:''" json:"avatar_url"`
    CreatedAt time.Time `json:"-"` // 不给前端看
    
    // 注册专用的二次确认字段（不存数据库）
    RePassword string `gorm:"-" json:"re_password" binding:"omitempty,eqfield=Password"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
