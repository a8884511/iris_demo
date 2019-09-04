package main

import (
	"Goo/app"
	"github.com/kataras/iris"
)

func main() {
	application, onExit := app.NewApp()
	defer onExit()
	application.Run(iris.Addr("localhost:10001"))
}
