package util

import (
	"Goo/app/db"
	"Goo/app/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/juju/errors"
	"github.com/kataras/iris"
	"time"
)

var (
	JWTSecretKey       = []byte("you-will-never-guess")
	JWTSecretKeyGetter = func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	}
)

func GenerateJWToken(user model.User) (string, error) {
	claim := jwt.MapClaims{
		"userid":   user.ID,
		"username": user.Username,
		//Token签发者，格式是区分大小写的字符串或者uri，用于唯一标识签发token的一方。
		"iss": "Datagrand",
		//Token的主体，即它的所有人，格式是区分大小写的字符串或者uri。
		"sub": "Anyone who is a datagrand user",
		//指定Token在nbf时间之前不能使用，即token开始生效的时间，格式为时间戳。
		"nbf": time.Now().Unix(),
		//Token的签发时间，格式为时间戳。
		"iat": time.Now().Unix(),
		//Token的过期时间，格式为时间戳。
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(JWTSecretKey)
	return tokenStr, err
}

func ParseJWToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, JWTSecretKeyGetter)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		err = errors.New("token is invalid")
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), err
}

func GetCurrentClaims(ctx iris.Context) (jwt.MapClaims, error) {
	token := ctx.Values().Get("jwt")
	if token == nil {
		return nil, errors.New("cannot get jwt claims")
	}
	return token.(*jwt.Token).Claims.(jwt.MapClaims), nil
}

func GetCurrentUser(ctx iris.Context) (user model.User, err error) {
	claims, err := GetCurrentClaims(ctx)
	if err != nil {
		return
	}
	result := db.Session.First(&user, "id = ?", claims["userid"])
	if *(user.IsActive) != true {
		err = errors.New("User is not active")
		return
	}
	return user, result.Error
}
