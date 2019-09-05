package admin

import (
	"Goo/app/db"
	"Goo/app/model"
	"Goo/app/util"
	"github.com/kataras/iris"
)

//plugin.Enforcer.AddPolicy("superuser", "/api/protect/jwt_required2", "GET")
//plugin.Enforcer.AddRoleForUser("admin", "superuser")
//plugin.Enforcer.LoadPolicy()
//db.Session.LogMode(true) //show sql

//get list
func GroupListApi(ctx iris.Context) {
	var groups []model.Group
	var count uint
	pagination, err := util.GetPagination(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if result := db.Session.Offset(pagination.Offset).Limit(pagination.Limit).Find(&groups).Offset(-1).Limit(-1).Count(&count); result.Error != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "An error occurred while querying"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"result": groups,
		"count":  count,
	})
}

//post list
func GroupCreateApi(ctx iris.Context) {

}

//get item
func GroupDetailApi(ctx iris.Context) {

}

//put/patch item
func GroupUpdateApi(ctx iris.Context) {

}

//delete
func GroupDeleteApi(ctx iris.Context) {

}
