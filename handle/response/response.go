package response

import (
	"KpChatGpt/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Data 数据返回结构
type Data struct {
	Code int         `json:"code"`           // 状态码
	Msg  string      `json:"msg"`            // 状态描述
	Data interface{} `json:"data,omitempty"` // 常规json数据或PaginationData数据
}

// MessageData 消息数据返回结构
type MessageData struct {
	Code uint32 `json:"code"` // 状态码
	Msg  string `json:"msg"`  // 状态描述
}

// JsonMessage 消息返回
func JsonMessage(ctx *gin.Context, code int, msg string) {
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, Data{
		Code: code,
		Msg:  msg,
	})
	ctx.Abort()
}

// JsonSuccess 成功带结果返回
func JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.Status(http.StatusOK)
	ctx.JSON(200, Data{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
	ctx.Abort()
}

// JsonFailMessage 失败信息返回
func JsonFailMessage(ctx *gin.Context, code int, err error) {
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, Data{
		Code: code,
		Msg:  errorss.HandleError(code, "zn", err).Error(),
	})
	ctx.Abort()
}
