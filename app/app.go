package app

import (
	"Goo/app/api"
	"Goo/app/db"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func NewApp() (*iris.Application, func()) {
	app := iris.New()
	app.Use(recover.New())

	app.Logger().SetLevel("debug")
	app.Use(consoleLogger)

	fileLogger, fileLoggerCloseFunc := newFileLogger()
	app.Use(fileLogger)

	db.Connect("sqlite3", "db.sqlite")
	db.Migrate()

	app.PartyFunc("/api", api.RegisterUrls)

	onExit := func() {
		fileLoggerCloseFunc()
	}
	return app, onExit
}
