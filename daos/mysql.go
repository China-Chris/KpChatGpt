package daos

import (
	"KpChatGpt/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db 全局MySQL数据库操作对象
var db *gorm.DB

//InitMysql 链接数据库
func InitMysql() {
	cfg := configs.Config().Mysql
	if cfg.Ip == "" {
		panic("invalid mysql ip")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Ip, cfg.Port, cfg.DbName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
