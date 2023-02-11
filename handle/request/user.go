package request

// ReLogin 登录请求参数
type ReLogin struct {
	Phone   string `form:"phone"`    //手机号
	SmsCode string `form:"sms_code"` //短信
	AliAuth string `form:"ali_auth"` //支付宝权限
	WxAuth  string `form:"wx_auth"`  //微信权限
}

// RqSignUp 注册请求参数
type RqSignUp struct {
	Phone    string `json:"phone"`    //手机号
	Password string `json:"password"` //密码
	Status   int    `json:"status"`   //注册状态(1.手机号密码注册)
}
