package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	//auth info
	Username string `gorm:"UNIQUE; NOT NULL"`
	//TODO 如何定义私有变量且创建表字段
	PasswordHash string
	IsActive     bool
	//other info
	Nickname string
	Birthday *time.Time
}

func (u User) TableName() string {
	return "user"
}

func (u *User) SetPassword(rawPassword string) {
	u.PasswordHash = MakePassword(rawPassword)
}

func (u User) CheckPassword(rawPassword string) bool {
	return MakePassword(rawPassword) == u.PasswordHash
}

func MakePassword(rawPassword string) string {
	//TODO 密码哈希值+盐值
	return rawPassword
}
