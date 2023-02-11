package main

import (
	"KpChatGpt/cache"
	"KpChatGpt/configs"
	"KpChatGpt/controller/gpt3"
	"KpChatGpt/daos"
	logger "KpChatGpt/logs"
	"KpChatGpt/routers"
	"KpChatGpt/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ParseConfig("./configs/cfg.yml")                                   //配置读取
	services.InitClient("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") //默认单key
	daos.InitMysql()                                                           //初始化数据库
	cache.NewRedis()                                                           //初始化缓存
	logger.InitLog()                                                           //日志
}

func main() {
	go gpt3.Manager.Start() //开启gpt3-监听
	r := gin.Default()
	routers.Route(r)
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8888) //拼接监听地址
	r.Run(address)
}
