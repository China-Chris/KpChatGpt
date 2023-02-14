package users

import (
	"KpChatGpt/cache"
	"KpChatGpt/e/errors_const"
	"KpChatGpt/handle/request"
	"KpChatGpt/handle/response"
	"KpChatGpt/services/user"
	"github.com/gin-gonic/gin"
)

// Login 用户登陆
func Login(ctx *gin.Context) {
	var loginReq request.RqLogin
	var data string
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		response.JsonFailMessage(ctx, errors_const.ErrInternalServer, err) //json解析失败
		return
	}
	//checkMobile := user.CheckMobile(signUp.Phone) //检查手机号
	//if !checkMobile {
	//	response.JsonFailMessage(ctx, errors_const.ErrCheckMobile, err)
	//	return
	//}
	if loginReq.Phone != "" && loginReq.SmsCode != "" { // 手机号登录
		checkSms, err := cache.VerifyCodeFromRedis(loginReq.Phone, loginReq.SmsCode)
		if err != nil {
			response.JsonFailMessage(ctx, errors_const.ErrCheckSms, err) //json解析失败
			return
		}
		data, err = user.FindUserByPhone(loginReq.Phone, loginReq.SmsCode, checkSms)
		if err != nil {
			response.JsonFailMessage(ctx, errors_const.ErrCheckSms, err) //json解析失败
			return
		}
	}
	//// 支付宝登录
	//if loginReq.AliAuth != "" {
	//	user, err := findUserByAliAuth(loginReq.AliAuth)
	//	if err != nil {
	//		response.JsonBadRequest(ctx, err.Error())
	//		return
	//	}
	//
	//	// 用户验证通过
	//	// ...
	//}
	//// 微信登录
	//if loginReq.WxAuth != "" {
	//	user, err := findUserByWxAuth(loginReq.WxAuth)
	//	if err != nil {
	//		response.JsonBadRequest(ctx, err.Error())
	//		return
	//	}
	//
	//	// 用户验证通过
	//	// ...
	//}
	response.JsonSuccess(ctx, data)
}

// Sms 获取sms短信
func Sms(ctx *gin.Context) {
	var signUp request.RqSms
	err := ctx.ShouldBindJSON(&signUp)
	if err != nil {
		response.JsonFailMessage(ctx, errors_const.ErrInternalServer, err) //json解析失败
		return
	}
	//checkMobile := user.CheckMobile(signUp.Phone) //检查手机号
	//if !checkMobile {
	//	response.JsonFailMessage(ctx, errors_const.ErrCheckMobile, err)
	//	return
	//}
	err = user.GetSms(signUp.Phone) //获取SMS短信
	if err != nil {
		response.JsonFailMessage(ctx, errors_const.ErrSendSMS, err)
		return
	}
	response.JsonSuccess(ctx, nil)
}
