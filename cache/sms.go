package cache

import (
	"KpChatGpt/configs"
	errorss "KpChatGpt/e"
	"KpChatGpt/e/errors_const"
	"fmt"
	"time"
)

// SaveCodeToRedis 将验证码存储到redis中
func SaveCodeToRedis(phone, code string) error {
	key := fmt.Sprintf("sms_%s", phone)
	err := Rdb.Set(configs.Ctx, key, code, time.Minute*5).Err()
	if err != nil {
		return errorss.HandleError(errors_const.ErrSaveSmsToRedis, "zn", err)
	}
	return nil
}

// VerifyCodeFromRedis 从redis中验证验证码
func VerifyCodeFromRedis(phone, code string) (bool, error) {
	key := fmt.Sprintf("sms_%s", phone)
	c, err := Rdb.Get(configs.Ctx, key).Result()
	if err != nil {
		return false, errorss.HandleError(errors_const.ErrVerifySmsFromRedis, "zn", err)
	}
	if c != code {
		return false, errorss.HandleError(errors_const.ErrInvalidSmsCode, "zn", nil)
	}
	return false, nil
}
