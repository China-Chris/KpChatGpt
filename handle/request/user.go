package request

// RqLogin 登录请求参数
type RqLogin struct {
	Phone   string `form:"phone"`    //手机号
	SmsCode string `form:"sms_code"` //短信
	AliAuth string `form:"ali_auth"` //支付宝权限
	WxAuth  string `form:"wx_auth"`  //微信权限
}

// RqSms 短信请求参数
type RqSms struct {
	Phone string `json:"phone"` //手机号
}
