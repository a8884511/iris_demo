package app

import (
	"Goo/app/api"
	"Goo/app/db"
	"Goo/app/plugin"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func NewApp() (*iris.Application, func()) {
	app := iris.New()

	//recover from panic
	app.Use(recover.New())

	//logger
	app.Logger().SetLevel("debug")
	app.Use(consoleLogger)
	fileLogger, fileLoggerCloseFunc := newFileLogger()
	app.Use(fileLogger)

	//database
	db.Connect("sqlite3", "db.sqlite")
	db.Migrate()

	//plugins
	plugin.CasbinInitialize()

	//routes
	app.PartyFunc("/api", api.RegisterUrls)

	//error handlers
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, InternalServerError)

	onExit := func() {
		fileLoggerCloseFunc()
	}
	return app, onExit
}
