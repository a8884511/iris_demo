package greet

import (
	"github.com/kataras/iris"
)

func RegisterUrls(GreetParty iris.Party) {
	GreetParty.Use()
	{
		GreetParty.Get("/hello", HelloApi)
	}
}
