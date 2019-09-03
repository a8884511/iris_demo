package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(rawPassword)); err != nil {
		return false
	}
	return true
}

func MakePassword(rawPassword string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password")
	}
	return string(passwordHash)
}
