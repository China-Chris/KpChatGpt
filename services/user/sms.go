package user

import (
	"KpChatGpt/cache"
	"KpChatGpt/configs"
	errorss "KpChatGpt/e"
	"KpChatGpt/e/errors_const"
	"fmt"
	"github.com/recallsong/httpc"
	"github.com/zhengjianfeng1103/vUtils/random"
)

const (
	Message      = "【chat】验证码：%s，有效期为5分钟，如非本人操作，请忽略本短信"
	UserID       = "66654"
	RandomLength = 6
)

// SendSmsRequest 请求参数结构体
type SendSmsRequest struct {
	Action      string `url:"action"`
	UserID      string `url:"userid"`
	Account     string `url:"account"`
	Password    string `url:"password"`
	Mobile      string `url:"mobile"`
	Content     string `url:"content"`
	SendTime    string `url:"sendTime"`
	ExtensionNo string `url:"extno"`
}

// ReMessage 短信发送返回信息结构体
type ReMessage struct {
	Returnstatus  string
	Message       string
	Remainpoint   string
	TaskID        string
	SuccessCounts string
}

// GetSms 获取Sms短信
func GetSms(phone string) error {
	cfg := configs.Config().Message
	var messages ReMessage
	code, err := random.GenerateSafeString(RandomLength)
	if err != nil {
		return errorss.HandleError(errors_const.ErrGenerateToken, "zn", err)
	}
	content := fmt.Sprintf(Message, code)
	err = httpc.New(cfg.Message).Path(cfg.MessagePath).
		Query("action", "send").
		Query("userid", UserID).
		Query("account", "OA00321").
		Query("password", "39B6C57B96DA623BD95207C205595CF6").
		Query("mobile", phone).
		Query("content", content).
		Query("sendTime", "").
		Query("extno", "").
		Get(&messages)
	if err != nil {
		return errorss.HandleError(errors_const.ErrSendSMS, "zn", err)
	}
	if messages.Returnstatus != "Success" || messages.SuccessCounts != "1" {
		return errorss.HandleError(errors_const.ErrSendSMS, "zn", err)
	}
	err = cache.SaveCodeToRedis(phone, code)
	if err != nil {
		return errorss.HandleError(errors_const.ErrSendSMS, "zn", err)
	}
	return nil
}
