package routers

import (
	"KpChatGpt/handle"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.Use(handle.Cors())
	group := r.Group("/api/v1")
	{
		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		group.GET("/gep3", handle.Gpt)
	}
}
