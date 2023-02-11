package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户表
type User struct {
	gorm.Model
	Phone    string    `gorm:"type:varchar(20);not null"`  //手机号
	Password string    `gorm:"type:varchar(50);not null"`  //密码
	UserName string    `gorm:"type:varchar(50);not null"`  //用户名
	Email    string    `gorm:"type:varchar(100);not null"` //邮箱
	NickName string    `gorm:"type:varchar(50)"`           //用户昵称
	Avatar   string    `gorm:"type:varchar(255);not null"` //头像
	Age      int       //年龄
	Gender   int       `gorm:"not null"` //性别
	Birthday time.Time //生日
	Address  string    `gorm:"type:varchar(255);not null"` //地址
	Status   int       `gorm:"not null"`                   //用户状态(是否激活，是否禁用等)
	Role     int       `gorm:"not null"`                   //角色
}

// LoginRecord 登录记录
type LoginRecord struct {
	gorm.Model
	UserId    uint      `gorm:"not null"` //用户表ID
	LoginTime time.Time `gorm:"not null"` //最后登陆时间
	Age       int       //年龄
	Gender    string    //性别
	Location  string    //位置
}

// UserActivity 统计用户活动
type UserActivity struct {
	Date          time.Time //日期
	DailyActive   int       //日活
	MonthlyActive int       //月活
	RetentionRate float64   //留存率
}

func (t User) TableName() string {
	return "user"
}
func (t LoginRecord) TableName() string {
	return "login_record"
}

func (t UserActivity) TableName() string {
	return "user_activity"
}
