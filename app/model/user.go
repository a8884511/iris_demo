package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Username     string `gorm:"UNIQUE; NOT NULL"`
	PasswordHash string
	IsActive     bool
	Nickname     string
	Birthday     *time.Time
}

func (u User) TableName() string {
	return "user"
}

func (u *User) SetPassword(rawPassword string) error {
	passwordHash, err := HashPassword(rawPassword)
	if err != nil {
		return err
	}
	u.PasswordHash = passwordHash
	return nil
}

func (u User) CheckPassword(rawPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(rawPassword)); err != nil {
		return err
	}
	return nil
}

func HashPassword(rawPassword string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(passwordHash), err
}
