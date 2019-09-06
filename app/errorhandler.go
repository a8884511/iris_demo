package app

import "github.com/kataras/iris"

func NotFound(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Not Found"})
}

func InternalServerError(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Internal Server Error"})
}
