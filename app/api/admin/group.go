package admin

import (
	"Goo/app/db"
	"Goo/app/model"
	"Goo/app/util"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
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
	if result := db.Session.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Users").Find(&groups).Offset(-1).Limit(-1).Count(&count); result.Error != nil {
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

type GroupForm struct {
	Name string `json:"name" validate:"required"`
}

//post list
func GroupCreateApi(ctx iris.Context) {
	validate := validator.New()
	var form GroupForm
	if err := ctx.ReadJSON(&form); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := validate.Struct(form); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"message": err.Error()})
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	var group model.Group
	if result := db.Session.First(&group, "name = ?", form.Name); result.Error == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Group name already exists"})
		return
	}
	group = model.Group{Name: form.Name}
	db.Session.Create(&group)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Created"})
}

//get item
func GroupDetailApi(ctx iris.Context) {
	var group model.Group
	if result := db.Session.Preload("Users").First(&group, "id = ?", ctx.Params().Get("id")); result.Error != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": "Not Found"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(group)
}

//put/patch item
func GroupUpdateApi(ctx iris.Context) {

}

//delete
func GroupDeleteApi(ctx iris.Context) {

}
