package cache

import (
	"KpChatGpt/configs"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client //配置全局redis客户端

// NewRedis 初始化Redis数据库
func NewRedis() {
	cfg := configs.Config().Redis
	Ctx := configs.Ctx
	nds := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client := redis.Options{
		Addr:     nds,
		Password: cfg.Passwd,
		DB:       cfg.DbName,
	}
	// 创建连接池
	Rdb = redis.NewClient(&client)
	pong, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		logs.Error("连接redis失败,%s，err:%s", pong, err)
		panic(err.Error())
	}
	logs.Info("Redis数据库初始化连接成功")
}
