package db

import (
	"Goo/app/model"
	"github.com/jinzhu/gorm"
)

var Session *gorm.DB

func Connect(t string, u string) *gorm.DB {
	var err error
	Session, err = gorm.Open(t, u)
	if err != nil {
		panic("Failed to connect database")
	}
	return Session
}

func Migrate() {
	if Session == nil {
		panic("Must connect database first")
	}
	Session.AutoMigrate(&model.User{})
}
