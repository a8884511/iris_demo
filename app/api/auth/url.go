package auth

import (
	"github.com/kataras/iris"
)

func RegisterUrls(party iris.Party) {
	party.Use()
	{
		party.Post("/sign_in", SignInApi)
		//不允许注册
		party.Post("/sign_up", SignUpApi)
		party.Post("/token_check", TokenCheckApi)
	}
}
