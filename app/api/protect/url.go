package protect

import (
	"Goo/app/middleware"
	"Goo/app/plugin"
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
	plugin.Enforcer.AddPolicy("superuser", "/api/protect/jwt_required2", "GET")
	plugin.Enforcer.AddRoleForUser("admin", "superuser")
	plugin.Enforcer.LoadPolicy()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"message": "Hello World",
		"user":    user,
		"rules":   plugin.Enforcer.GetAllObjects(),
	})
}

func JWTRequired2Api(ctx iris.Context) {
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

func RegisterUrls(ProtectParty iris.Party) {
	ProtectParty.Use(middleware.JWTMiddleware)
	{
		ProtectParty.Get("/jwt_required", JWTRequiredApi)
		ProtectParty.Get("/jwt_required2", middleware.CasbinMiddleware, JWTRequired2Api)
	}
}
