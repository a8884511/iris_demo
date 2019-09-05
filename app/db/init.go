package db

import (
	"Goo/app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
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
	Session.AutoMigrate(&model.Group{})
	Session.AutoMigrate(&model.User{})
	Session.AutoMigrate(&model.Role{})
	Session.AutoMigrate(&model.Api{})
}
