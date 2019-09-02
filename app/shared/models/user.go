package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	//auth info
	Username string `gorm:"UNIQUE; NOT NULL"`
	Password string
	IsActive bool
	//other info
	Nickname string
	Birthday *time.Time
}

func (u User) TableName() string {
	return "user"
}
