package request

// RqSignUp 注册请求参数
type RqSignUp struct {
	Phone    string `json:"phone"`    // 手机号
	Password string `json:"password"` //密码
	Status   int    `json:"status"`   //注册状态(1.手机号密码注册 2.验证码注册 3.第三方注册)
}
