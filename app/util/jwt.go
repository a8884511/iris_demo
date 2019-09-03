package util

import (
	"Goo/app/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/juju/errors"
	"time"
)

var JWTKey = []byte("you-will-never-guess")

func GenerateJWToken(user model.User) (string, error) {
	claim := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//TODO 配置管理
	tokenStr, err := token.SignedString(JWTKey)
	return tokenStr, err
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	}
}

func ParseJWToken(tokenStr string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenStr, secret())
	if err != nil {
		err = errors.New("Cannot parse token")
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("Cannot convert claim")
		return nil, err
	}
	if !token.Valid {
		err = errors.New("Token is invalid")
		return nil, err
	}
	return claim, err
}
