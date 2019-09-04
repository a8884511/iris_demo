package protect

import (
	"Goo/app/middleware"
	"github.com/kataras/iris"
)

func JWTRequiredApi(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Hello World"})
}

func RegisterUrls(ProtectParty iris.Party) {
	ProtectParty.Use(middleware.JWTMiddleware, middleware.ACLMiddleware)
	{
		ProtectParty.Get("/jwt_required", JWTRequiredApi)
	}
}
