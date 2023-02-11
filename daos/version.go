package daos

import (
	"KpChatGpt/model"
)

// GetVersionDao 获得版本Dao
func GetVersionDao() (version model.Version, err error) {
	if err = db.Model(model.Version{}).Where("id = ?", 1).Find(&version).Error; err != nil {
	}
	return
}
