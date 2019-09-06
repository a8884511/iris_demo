package app

import "github.com/kataras/iris"

func NotFound(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "not found"})
}

func InternalServerError(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "internal server error"})
}
