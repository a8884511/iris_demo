package main

import (
	. "Goo/app"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
)

func main() {
	app := NewApp()
	app.Run(iris.Addr("localhost:10001"))
}
