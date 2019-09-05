package greet

import (
	"github.com/kataras/iris"
)

func RegisterUrls(party iris.Party) {
	party.Use()
	{
		party.Get("/hello", HelloApi)
	}
}
