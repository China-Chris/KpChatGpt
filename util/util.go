package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RspMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseFront ...
// 封装的向前端返回响应的函数
func ResponseFront(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, RspMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
