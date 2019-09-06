package admin

import (
	"github.com/kataras/iris"
)

func RegisterUrls(party iris.Party) {
	party.Use()
	{
		party.Get("/groups", GroupListApi)
		party.Post("/groups", GroupCreateApi)
		party.Get("/groups/{id:uint}", GroupDetailApi)
	}
}
