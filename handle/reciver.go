package handle

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/gin-gonic/gin"
)

// AppRecover 全局异常捕获中间件
func AppRecover(ctx *gin.Context) {
	// 全局捕获异常
	defer func() {
		if err := recover(); err != nil {
			logs.Error("服务器发生异常错误：%s", fmt.Sprint(err))
			//tools.JsonFailed(ctx, 500, "网络异常")
		}
	}()
	// 调用下级
	ctx.Next()
}
