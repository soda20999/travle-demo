package vo

type ParamSignup struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamUpdateNickname struct {
    UserID   int64  `json:"user_id" binding:"required"`
    Nickname string `json:"nickname" binding:"required,max=64"`
}

type ParamUpdateAvatar struct {
    UserID   int64  `json:"user_id" binding:"required"`
    AvatarURL string `json:"avatar_url" binding:"required,url,max=256"`
}