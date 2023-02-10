package user

import (
	"KpChatGpt/e"
	"KpChatGpt/e/errors_const"
	"KpChatGpt/handle/request"
	"KpChatGpt/handle/response"
	"KpChatGpt/middleware/jwt"
	"github.com/beego/beego/v2/adapter/logs"
	"regexp"
)

// CheckMobile 检查手机号
func CheckMobile(phone string) bool {
	compile, err := regexp.Compile("^(?:\\+?86)?1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[0-35-9]\\d{2}|4(?:0\\d|" +
		"1[0-2]|9\\d))|9[0-35-9]\\d{2}|6[2567]\\d{2}|4(?:(?:10|4[01])\\d{3}|[68]\\d{4}|[579]\\d{2}))\\d{6}$")
	if err != nil {
		panic(err)
	}
	return !compile.MatchString(phone)
}

// SignUp 用户注册
func SignUp(that request.RqSignUp) (*response.RpSignUp, error) {
	logs.Info(that.Phone)
	accessToken, refreshToken, err := jwt.GenerateToken(that.Phone, that.Password)
	if err != nil {
		return nil, errorss.HandleError(errors_const.ErrGenerateToken, "zn", err)
	}
	return response.NewRpSignUp(accessToken, refreshToken), nil
}
