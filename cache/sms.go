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
