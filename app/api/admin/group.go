package admin

import (
	"Goo/app/db"
	"Goo/app/model"
	"Goo/app/util"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v9"
)

func GetGroupItemByID(id uint) (group model.Group, err error) {
	result := db.Session.Preload("Users").Preload("Roles").Preload("Apis").First(&group, "id = ?", id)
	return group, result.Error
}

func GetGroupItemByName(name string) (group model.Group, err error) {
	result := db.Session.Preload("Users").Preload("Roles").Preload("Apis").First(&group, "name = ?", name)
	return group, result.Error
}

func GetGroupList(pagination util.Pagination) (groups []model.Group, count uint, err error) {
	result := db.Session.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Users").Preload("Roles").Preload("Apis").Find(&groups).Offset(-1).Limit(-1).Count(&count)
	return groups, count, result.Error
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

type GroupCreateForm struct {
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc"`
}

func GroupCreateApi(ctx iris.Context) {
	validate := validator.New()
	var form GroupCreateForm
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
	if err := mapstructure.Decode(form, &group); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	group.UpdateStatus(map[string]interface{}{
		"desc": form.Desc,
	})
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

type GroupUpdateForm struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func GroupUpdateApi(ctx iris.Context) {
	validate := validator.New()
	var form GroupUpdateForm
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
	if form.Name != "" {
		g, err := GetGroupItemByName(form.Name)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON(iris.Map{"message": err.Error()})
				return
			}
		} else if g.Name != group.Name {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "group name already exists"})
			return
		}
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
