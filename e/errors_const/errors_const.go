package errors_const

// 用于存放各模块定义的错误常量，如过业务错误太多可以单独将各模块分成多个文件存放 例：user_const.go 存放 用户模块定义的错误

const ( //定义常规错误
	ErrInternalServer = iota + 500 //Post 解析失败 参数错误
)

const ( //  用户模块
	ErrShouldBind = iota + 1000 //Post 解析失败 参数错误
	ErrSignUp
	ErrCheckMobile
	ErrGenerateToken
	ErrGenerateSafeString
	ErrSendSMS
	ErrSaveSmsToRedis
	ErrVerifySmsFromRedis
	ErrInvalidSmsCode
	ErrCheckSms
	ErrFindUserByPhone
)
