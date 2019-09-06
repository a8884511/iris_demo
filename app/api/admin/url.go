package admin

import (
	"github.com/kataras/iris"
)

func RegisterUrls(party iris.Party) {
	party.Use()
	{
		//group
		party.Get("/groups", GroupListApi)
		party.Post("/groups", GroupCreateApi)
		party.Get("/groups/{id:uint}", GroupDetailApi)
		party.Patch("/groups/{id:uint}", GroupUpdateApi)
		party.Delete("/groups/{id:uint}", GroupDeleteApi)
		//user
		party.Get("/users", UserListApi)
		party.Post("/users", UserCreateApi)
		party.Get("/users/{id:uint}", UserDetailApi)
		party.Patch("/users/{id:uint}", UserUpdateApi)
		party.Delete("/users/{id:uint}", UserDeleteApi)
		//role
		party.Get("/roles", RoleListApi)
		party.Post("/roles", RoleCreateApi)
		party.Get("/roles/{id:uint}", RoleDetailApi)
		party.Patch("/roles/{id:uint}", RoleUpdateApi)
		party.Delete("/roles/{id:uint}", RoleDeleteApi)
		//api
		party.Get("/apis", ApiListApi)
		party.Post("/apis", ApiCreateApi)
		party.Get("/apis/{id:uint}", ApiDetailApi)
		party.Patch("/apis/{id:uint}", ApiUpdateApi)
		party.Delete("/apis/{id:uint}", ApiDeleteApi)
	}
}
