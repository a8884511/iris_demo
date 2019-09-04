package middleware

import (
	"Goo/app/util"
	"github.com/dgrijalva/jwt-go"
	jwtMiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

var JWTMiddleware = func(ctx iris.Context) {
	jwtMiddleware.New(jwtMiddleware.Config{
		ValidationKeyGetter: util.JWTSecretKeyGetter,
		ErrorHandler:        JWTErrorHandler,
		SigningMethod:       jwt.SigningMethodHS256,
	}).Serve(ctx)
}

func JWTErrorHandler(ctx iris.Context, err error) {
	ctx.StatusCode(iris.StatusForbidden)
	ctx.JSON(iris.Map{"message": err.Error()})
}
