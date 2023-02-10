package response

// RpSignUp 注册返回参数
type RpSignUp struct {
	accessToken  string `json:"accessToken"`
	refreshToken string `json:"refreshToken"`
}

// NewRpSignUp 返回注册参数
func NewRpSignUp(accessToken, refreshToken string) *RpSignUp {
	return &RpSignUp{accessToken, refreshToken}
}
