package app

import (
	"iam/internal/business/user/domin/dto"
	"iam/internal/business/user/domin/vo"
	"iam/internal/business/user/repo"
	"iam/internal/pkg/jwt"
	"iam/pkg/snowflake"
)

func SignUp(p *vo.ParamSignup) (Error error) {
	//判断用户存不存在
	boo, err := repo.CheckUserExists(p.Username)
	
	if err != nil {
		//return err //数据库错误
	}
	if boo {
		//return repo.ErrorUserExist
	}

	//生成UID
	userId := snowflake.GenID()
	user := dto.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}

	//数据存入数据库
	return repo.InsertUser(user)
}

func Login(p *vo.ParamLogin) (atoken, rToken string, err error) {
	user := dto.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = repo.Login(&user); err != nil {
		return "", "", err
	}

	return jwt.GenToken(user.Username, user.UserID)
}


// 更新用户昵称
func UpdateNickname(p *vo.ParamUpdateNickname) error {
    // 验证用户是否存在
    _, err := repo.GetUserByID(p.UserID)
    if err != nil {
        return err
    }
    
    // 更新昵称
    return repo.UpdateUserNickname(p.UserID, p.Nickname)
}

// 更新用户头像
func UpdateAvatar(p *vo.ParamUpdateAvatar) error {
    // 验证用户是否存在
    _, err := repo.GetUserByID(p.UserID)
    if err != nil {
        return err
    }
    
    // 更新头像
    return repo.UpdateUserAvatar(p.UserID, p.AvatarURL)
}

// 获取用户信息
func GetUserInfo(userID int64) (*dto.User, error) {
    user, err := repo.GetUserByID(userID)
    if err != nil {
        return nil, err
    }
    
    return dto.FromModel(user), nil
}
