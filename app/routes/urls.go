package routes

import (
	"Goo/app/routes/auth"
	"Goo/app/routes/greet"
	"github.com/kataras/iris"
)

func RegisterUrls(RootParty iris.Party) {
	GreetParty := RootParty.Party("/greet")
	greet.RegisterUrls(GreetParty)

	AuthParty := RootParty.Party("/auth")
	auth.RegisterUrls(AuthParty)
}
