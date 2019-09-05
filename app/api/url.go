package api

import (
	"Goo/app/api/admin"
	"Goo/app/api/auth"
	"Goo/app/api/greet"
	"Goo/app/api/protect"
	"github.com/kataras/iris"
)

func RegisterUrls(RootParty iris.Party) {
	RootParty.Use()
	{
		RootParty.PartyFunc("/greet", greet.RegisterUrls)
		RootParty.PartyFunc("/auth", auth.RegisterUrls)
		RootParty.PartyFunc("/protect", protect.RegisterUrls)
		RootParty.PartyFunc("/admin", admin.RegisterUrls)
	}
}
