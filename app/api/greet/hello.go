package greet

import (
	"github.com/kataras/iris"
)

func HelloApi(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Hello World"})
}
