package versions

import (
	"KpChatGpt/handle/response"
	"KpChatGpt/services/version"
	"github.com/gin-gonic/gin"
)

// GetVersion 获取版本
func GetVersion(ctx *gin.Context) {
	_, err := version.GetVersion()
	if err != nil {

	}
	response.JsonSuccess(ctx, nil)
}
