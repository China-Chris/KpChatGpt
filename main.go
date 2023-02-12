package main

import (
	"KpChatGpt/cache"
	"KpChatGpt/configs"
	"KpChatGpt/controller/gpt3"
	"KpChatGpt/daos"
	logger "KpChatGpt/logs"
	"KpChatGpt/routers"
	"KpChatGpt/services"
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var r *gin.Engine
var srv *http.Server

func init() {
	configs.ParseConfig("./configs/cfg.yml")                                   //配置读取
	services.InitClient("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") //默认单key
	daos.InitMysql()                                                           //初始化数据库
	cache.NewRedis()                                                           //初始化缓存
	logger.InitLog()                                                           //日志
}

func run() {
	go gpt3.Manager.Start() //开启gpt3-监听
	r = gin.Default()
	routers.Route(r)
	srv = &http.Server{
		Addr:    ":" + configs.Config().Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Error("Listen: ", err)
		}
	}()
}

func main() {
	go run()

	// 优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logs.Error("Server Shutdown:", err)
	}
	logs.Info("Server exiting")
}
