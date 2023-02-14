package routers

import (
	"KpChatGpt/controller/gpt3"
	"KpChatGpt/controller/users"
	"KpChatGpt/controller/versions"
	"KpChatGpt/handle"
	"KpChatGpt/middleware/activeCount"
	"KpChatGpt/middleware/bucket"
	"github.com/gin-gonic/gin"
)

const rate = 5 //令牌桶限制-每个ip每秒访问次数

func Route(r *gin.Engine) {
	r.Use(handle.AppRecover) //全局异常捕获
	r.Use(handle.Cors())     //全局跨域支持
	group := r.Group("/api/v1")
	group.Use(bucket.IpTokenBucketMiddleware(rate))
	{
		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success") //测试连通
		})
	}
	gpt := group.Group("/gpt")
	{
		gpt.GET("/gep3", gpt3.Gpt) //got3
	}
	user := group.Group("/user")
	{
		user.POST("/login", activeCount.DailyActiveCount, users.Login) //用户登陆
		user.POST("/sms", users.Sms)                                   //获取sms短信
		//user.POST("/captcha", users.captcha)                           //获取验证码
		//user.POST("/signUp",jwt.AuthMiddleware, activeCount.DailyActiveCount, users.SignUp)//编辑用户
		//user.DELETE("/signUp", activeCount.DailyActiveCount, users.SignUp)                 //用户注销
		//user.GET("/login", jwt.AuthMiddleware, users.Login)                                //用户邀请
	}
	version := group.Group("/version")
	{
		version.GET("/getVersion", versions.GetVersion) //获得版本
		//version.GET("/getVersion", versions.GetVersion) //新增版本
		//version.GET("/getVersion", versions.GetVersion) //更新版本
		//version.GET("/getVersion", versions.GetVersion) //删除版本
	}
}
