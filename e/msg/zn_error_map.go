package msg

import (
	"KpChatGpt/e/errors_const"
)

// ErrorCodeTextMapZh 定义中文的业务报错信息
var ErrorCodeTextMapZh = map[int]string{
	errors_const.ErrInternalServer:     "内部服务发生错误",
	errors_const.ErrShouldBind:         "参数解析错误请确认参数正确",
	errors_const.ErrSignUp:             "注册时发生错误",
	errors_const.ErrCheckMobile:        "手机号校验不匹配",
	errors_const.ErrGenerateToken:      "生成令牌发生错误",
	errors_const.ErrGenerateSafeString: "生成随机验证码错误",
	errors_const.ErrSendSMS:            "发送短信失败",
	errors_const.ErrSaveSmsToRedis:     "验证码存储到redis中发生错误",
	errors_const.ErrVerifySmsFromRedis: "验证验证码错误",
	errors_const.ErrInvalidSmsCode:     "验证失败,无效的验证码",
	errors_const.ErrCheckSms:           "短信验证失败，请重新获取",
	errors_const.ErrFindUserByPhone:    "账号登录失败",
}
