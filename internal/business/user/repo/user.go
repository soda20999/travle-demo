package repo

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"

	"iam/internal/business/user/domin/dto"
	"iam/internal/pkg/config/mysql"
)

const scretKey = "mumu.123.com"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// QueryUserByUserName 检查用户是否存在
func QueryUserByUserName(Username string) error {
	sqlString := "select count(user_id) from user where username=?"
	var count int
	if error := mysql.Db.Get(&count, sqlString, Username); error != nil {
		return error
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser 插入用户信息
func InsertUser(user dto.User) error {
	user.Password = encryptPassword(user.Password)

	sqlString := "insert into user(user_id,username,password) values(?,?,?)"
	_, err := mysql.Db.Exec(sqlString, user.UserID, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// Login 根据用户名和密码登录
func Login(user *dto.User) (err error) {
	opassowrd := user.Password
	sqlString := "select user_id,username,password from user where username=?"
	err = mysql.Db.Get(user, sqlString, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err //数据库错误
	}

	//验证密码是否错误
	password := encryptPassword(opassowrd)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}

func encryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(opassword + scretKey))
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}

