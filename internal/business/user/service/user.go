package user_service

import (
     "iam/internal/business/user/model"
	 "iam/internal/pkg/config/postsql"
	 "errors"
	 "golang.org/x/crypto/bcrypt"
	 "gorm.io/gorm"
)

// SignUp 注册：包含密码加密逻辑
func SignUp(u *model.User) error {
	// 1. 将明文密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	// 2. 写入 PostgreSQL (利用 GORM 的自动同步和唯一索引约束)
	return postgresql.DB.Create(u).Error
}

// Login 登录：比对密码并返回用户信息
func Login(username, password string) (*model.User, error) {
	var u model.User
	// 1. 根据用户名查找用户
	err := postgresql.DB.Where("username = ?", username).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 2. 比对哈希密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	return &u, nil
}

// UpdateUser 直接按 ID 更新任意字段 (利用 PG 的局部更新特性)
func UpdateUser(userID int64, data map[string]interface{}) error {
	return postgresql.DB.Model(&model.User{}).Where("user_id = ?", userID).Updates(data).Error
}

// GetUserInfo 获取信息 (去 DTO 化：直接返回 Model)
func GetUserInfo(userID int64) (u model.User, err error) {
	err = postgresql.DB.Where("user_id = ?", userID).First(&u).Error
	return
}