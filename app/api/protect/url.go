package protect

import (
	"Goo/app/middleware"
	"Goo/app/util"
	"github.com/kataras/iris"
)

func JWTRequiredApi(ctx iris.Context) {
	user, err := util.GetCurrentUser(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"message": "Hello World",
		"user":    user,
	})
}

func RegisterUrls(party iris.Party) {
	party.Use(middleware.JWTMiddleware)
	{
		party.Get("/jwt_required", JWTRequiredApi)
	}
}
