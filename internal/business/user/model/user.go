// d:/go/project/iam/internal/business/user/model/user.go
package model

import (
	"time"
)

type User struct {
	ID        int64    `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Username  string    `gorm:"column:username;size:255;not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"column:password;size:255;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}