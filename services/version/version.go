package version

import (
	"KpChatGpt/daos"
	"KpChatGpt/model"
)

// GetVersion 获取版本services
func GetVersion() (version model.Version, err error) {
	return daos.GetVersionDao()
}
