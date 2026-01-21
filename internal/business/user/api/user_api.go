package api

import (
	"errors"

	"iam/internal/business/user/app"
	"iam/internal/business/user/domin/vo"
	"iam/internal/business/user/repo"
	"iam/internal/pkg/response"
	"iam/internal/pkg/validator"

	"github.com/gin-gonic/gin"
	vd "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignHandler(c *gin.Context) {
	//1.获取参数
	p := new(vo.ParamSignup)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(&p) error failed :", zap.Error(err))
		//判断error是不是validator的错误类型
		errors, ok := err.(vd.ValidationErrors)
		if !ok {
			//不是validator的错误类型，说明是其他错误
			response.ResponseError(response.CodeInvalidParams, c)
			return
		}

		//是validator的错误类型，则进行翻译
		response.ResponseErrorWithMsg(validator.RemoveTopStruct(errors.Translate(validator.Trans)), c)
		return
	}

	//3.业务处理
	if err := app.SignUp(p); err != nil {
		zap.L().Error("SignUp error failed :", zap.Error(err))
		if errors.Is(err, repo.ErrorUserExist) {
			response.ResponseError(response.CodeUserExist, c)
			return
		}
		response.ResponseError(response.CodeServerBusy, c)
		return
	}

	//4.返回响应
	response.ResponseSuccess(nil, c)
}

func LoginHandler(c *gin.Context) {
	p := new(vo.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ShouldBindJSON error failed :", zap.Error(err))
		errors, ok := err.(vd.ValidationErrors)
		if !ok {
			response.ResponseError(response.CodeInvalidParams, c)
			return
		}
		response.ResponseErrorWithMsg(validator.RemoveTopStruct(errors.Translate(validator.Trans)), c)
		return
	}
	//3.业务处理
	atoken, _, err := app.Login(p)
	if err != nil {
		zap.L().Error("Login error failed :", zap.Error(err))
		if errors.Is(err, repo.ErrorUserNotExist) {
			response.ResponseError(response.CodeUserNotExist, c)
			return
		}
		response.ResponseError(response.CodeInvalidPassword, c)
		return
	}

	//4.返回响应
	response.ResponseSuccess(c, atoken)
}
