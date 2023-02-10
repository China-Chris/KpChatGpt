package main

import (
	configs "KpChatGpt/configs"
	"KpChatGpt/handle"
	logger "KpChatGpt/logs"
	"KpChatGpt/routers"
	"KpChatGpt/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ParseConfig("./configs/cfg.yml") //配置读取
	services.InitClient("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")
	//daos.InitMysql()                         //数据库
	//cache.NewRedis()                         //缓存
	logger.InitLog() //日志
}

func main() {
	go handle.Manager.Start()
	r := gin.Default()
	routers.Route(r)
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8888) //拼接监听地址
	r.Run(address)
}
