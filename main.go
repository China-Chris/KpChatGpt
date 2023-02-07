package main

import (
	"KpChatGpt/handle"
	"KpChatGpt/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//configs.ParseConfig("") //配置读取
	services.InitClient("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")

	go handle.Manager.Start()
	r := gin.Default()
	route(r)
	//r.Run(":" + configs.Config().Port)
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8888) //拼接监听地址
	r.Run(address)
}

func route(r *gin.Engine) {
	r.Use(handle.Cors())
	group := r.Group("/api/v1")
	{
		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		group.GET("/gepQa", handle.Gpt)
	}
}
