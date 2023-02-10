package routers

import (
	"KpChatGpt/controller/gpt3"
	"KpChatGpt/handle"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.Use(handle.Cors())
	group := r.Group("/api/v1")
	{
		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success") //测试连通
		})
		group.GET("/gep3", gpt3.Gpt) //got3

	}
	//user := r.Group("/api/v1/user")
	//{
	//	//user.GET("/login",)
	//}

}
