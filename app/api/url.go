package api

import (
	"Goo/app/api/admin"
	"Goo/app/api/auth"
	"Goo/app/api/greet"
	"Goo/app/api/protect"
	"github.com/kataras/iris"
)

func RegisterUrls(party iris.Party) {
	party.Use()
	{
		party.PartyFunc("/greet", greet.RegisterUrls)
		party.PartyFunc("/auth", auth.RegisterUrls)
		party.PartyFunc("/protect", protect.RegisterUrls)
		party.PartyFunc("/admin", admin.RegisterUrls)
	}
}
