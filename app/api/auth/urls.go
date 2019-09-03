package auth

import (
	"github.com/kataras/iris"
)

func RegisterUrls(AuthParty iris.Party) {
	AuthParty.Use()
	{
		AuthParty.Post("/sign_in", SignInApi)
		AuthParty.Post("/sign_up", SignUpApi)
	}
}
