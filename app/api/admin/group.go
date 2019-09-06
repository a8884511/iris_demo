package admin

import (
	"Goo/app/db"
	"Goo/app/model"
	"Goo/app/util"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

func GetGroupItemByID(id uint) (group model.Group, err error) {
	result := db.Session.Preload("Users").First(&group, "id = ?", id)
	return group, result.Error
}

func GetGroupItemByName(name string) (group model.Group, err error) {
	result := db.Session.Preload("Users").First(&group, "name = ?", name)
	return group, result.Error
}

func GetGroupList(pagination util.Pagination) (groups []model.Group, count uint, err error) {
	result := db.Session.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Users").Find(&groups).Offset(-1).Limit(-1).Count(&count)
	return groups, count, result.Error
}

type GroupForm struct {
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc"`
}

func GroupListApi(ctx iris.Context) {
	var groups []model.Group
	var count uint
	pagination, err := util.GetPagination(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	groups, count, err = GetGroupList(pagination)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"result": groups,
		"count":  count,
	})
}

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
	group, err := GetGroupItemByName(form.Name)
	if err == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "group name already exists"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	group = model.Group{Name: form.Name, BaseModel: model.BaseModel{Desc: form.Desc}}
	db.Session.Create(&group)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "created",
		"result":  group,
	})
}

func GroupDetailApi(ctx iris.Context) {
	var group model.Group
	id, err := ctx.Params().GetUint("id")
	group, err = GetGroupItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(group)
}

func GroupUpdateApi(ctx iris.Context) {
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
	id, err := ctx.Params().GetUint("id")
	group, err = GetGroupItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	db.Session.Model(&group).Update(form)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "updated",
		"result":  group,
	})
}

func GroupDeleteApi(ctx iris.Context) {
	var group model.Group
	id, err := ctx.Params().GetUint("id")
	group, err = GetGroupItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	db.Session.Delete(&group)
	ctx.StatusCode(iris.StatusNoContent)
}
