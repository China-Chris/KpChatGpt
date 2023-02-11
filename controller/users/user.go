package users

import (
	"KpChatGpt/e/errors_const"
	"KpChatGpt/handle/request"
	"KpChatGpt/handle/response"
	"KpChatGpt/services/user"
	"github.com/gin-gonic/gin"
)

// Login 用户登陆
func Login(ctx *gin.Context) {
	response.JsonSuccess(ctx, nil)
}

// SignUp 用户注册
func SignUp(ctx *gin.Context) {
	var signUp request.RqSignUp
	err := ctx.ShouldBindJSON(&signUp)
	if err != nil {
		response.JsonFailMessage(ctx, errors_const.ErrInternalServer, err) //json解析失败
		return
	}
	checkMobile := user.CheckMobile(signUp.Phone)
	if !checkMobile {
		response.JsonFailMessage(ctx, errors_const.ErrCheckMobile, err)
		return
	}
	date, err := user.SignUp(signUp)
	if err != nil {
		response.JsonFailMessage(ctx, errors_const.ErrSignUp, err)
		return
	}
	response.JsonSuccess(ctx, date)
}
