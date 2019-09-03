package api

import (
	"Goo/app/api/auth"
	"Goo/app/api/greet"
	"github.com/kataras/iris"
)

func RegisterUrls(RootParty iris.Party) {
	RootParty.Use()
	{
		RootParty.PartyFunc("/greet", greet.RegisterUrls)
		RootParty.PartyFunc("/auth", auth.RegisterUrls)
	}
}
