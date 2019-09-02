package auth

import (
	"github.com/kataras/iris"
)

func RegisterUrls(AuthParty iris.Party) {
	AuthParty.Use()
	{
		AuthParty.Post("/login", LoginApi)
	}
}
