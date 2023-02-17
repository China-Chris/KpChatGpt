package cache

import (
	"KpChatGpt/configs"
	"KpChatGpt/middleware/jwt"
	"time"
)

// cacheToken 将 Access Token 和 Refresh Token 放入 Redis 缓存
func cacheToken(accessToken, refreshToken string) error {
	// 设置 Access Token 过期时间
	shortExpireTime := time.Now().Add(jwt.Minute30)
	// 将 Access Token 放入 Redis 缓存中，设置过期时间为 30 分钟
	err := Rdb.Set(configs.Ctx, accessToken, "", shortExpireTime.Sub(time.Now())).Err()
	if err != nil {
		return err
	}
	// 设置 Refresh Token 过期时间
	longExpireTime := time.Now().Add(jwt.OneWeek)
	// 将 Refresh Token 放入 Redis 缓存中，设置过期时间为一周
	err = Rdb.Set(configs.Ctx, refreshToken, "", longExpireTime.Sub(time.Now())).Err()
	if err != nil {
		return err
	}
	return nil
}
