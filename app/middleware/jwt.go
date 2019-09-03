package middleware

import (
	"Goo/app/util"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: util.JWTSecretKeyGetter,
	ErrorHandler:        JWTErrorHandler,
	SigningMethod:       jwt.SigningMethodHS256,
})

func JWTErrorHandler(ctx iris.Context, err error) {
	ctx.JSON(iris.Map{"message": err.Error()})
}
