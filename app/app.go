package app

import (
	"Goo/app/api"
	"Goo/app/db"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func NewApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	Session := db.Connect("sqlite3", "db.sqlite")
	db.Migrate(Session)

	app.PartyFunc("/api", api.RegisterUrls)
	return app
}
