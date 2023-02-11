package model

import (
	"gorm.io/gorm"
	"time"
)

// Version 版本
type Version struct {
	gorm.Model
	VersionID   string    //版本Id
	IsMandatory bool      //是否强制更新
	ApkURL      string    //Apk地址
	Description string    //升级内容
	ReleasedAt  time.Time //版本发布时间
}

func (t Version) TableName() string {
	return "version"
}
