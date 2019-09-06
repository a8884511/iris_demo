package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	BaseModel
	Username     string     `gorm:"UNIQUE; NOT NULL" json:"username"`
	PasswordHash string     `json:"-"`
	IsActive     bool       `json:"is_active"`
	Nickname     string     `json:"nickname"`
	Birthday     *time.Time `json:"birthday"`
	GroupID      uint       `json:"group_id"`
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
