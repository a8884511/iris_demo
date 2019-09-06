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

func GetApiItemByID(id uint) (api model.Api, err error) {
	result := db.Session.Preload("Group").First(&api, "id = ?", id)
	return api, result.Error
}

func GetApiItemByMethodAndUrl(method string, url string) (api model.Api, err error) {
	result := db.Session.Preload("Group").First(&api, map[string]interface{}{
		"method": method,
		"url":    url,
	})
	return api, result.Error
}

func GetApiList(pagination util.Pagination) (apis []model.Api, count uint, err error) {
	result := db.Session.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Group").Find(&apis).Offset(-1).Limit(-1).Count(&count)
	return apis, count, result.Error
}

func ApiListApi(ctx iris.Context) {
	var apis []model.Api
	var count uint
	pagination, err := util.GetPagination(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	apis, count, err = GetApiList(pagination)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"result": apis,
		"count":  count,
	})
}

type ApiCreateForm struct {
	Method  string `json:"method" validate:"required"`
	Url     string `json:"url" validate:"required"`
	GroupID uint   `json:"group_id"`
	Desc    string `json:"desc"`
}

func ApiCreateApi(ctx iris.Context) {
	validate := validator.New()
	var form ApiCreateForm
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
	var api model.Api
	api, err := GetApiItemByMethodAndUrl(form.Method, form.Url)
	if err == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "api already exists"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := mapstructure.Decode(form, &api); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if form.GroupID != 0 {
		if _, err := GetGroupItemByID(form.GroupID); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": err.Error()})
			return
		}
	}
	api.UpdateStatus(map[string]interface{}{
		"desc": form.Desc,
	})
	db.Session.Create(&api)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "created",
		"result":  api,
	})
}

func ApiDetailApi(ctx iris.Context) {
	var api model.Api
	id, err := ctx.Params().GetUint("id")
	api, err = GetApiItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(api)
}

type ApiUpdateForm struct {
	Method  string `json:"method"`
	Url     string `json:"url"`
	GroupID uint   `json:"group_id"`
	Desc    string `json:"desc"`
}

func ApiUpdateApi(ctx iris.Context) {
	validate := validator.New()
	var form ApiUpdateForm
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
	var api model.Api
	id, err := ctx.Params().GetUint("id")
	api, err = GetApiItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if form.Method != "" && form.Url == "" {
		a, err := GetApiItemByMethodAndUrl(form.Method, api.Url)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON(iris.Map{"message": err.Error()})
				return
			}
		} else if a.Method != api.Method {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "api method already exists"})
			return
		}
	}
	if form.Method == "" && form.Url != "" {
		a, err := GetApiItemByMethodAndUrl(api.Method, form.Url)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON(iris.Map{"message": err.Error()})
				return
			}
		} else if a.Url != api.Url {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "api url already exists"})
			return
		}
	}
	if form.Method != "" && form.Url != "" {
		a, err := GetApiItemByMethodAndUrl(form.Method, form.Url)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON(iris.Map{"message": err.Error()})
				return
			}
		} else if a.Method != api.Method || a.Url != api.Url {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "api already exists"})
			return
		}
	}
	if form.GroupID != 0 {
		if _, err := GetGroupItemByID(form.GroupID); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": err.Error()})
			return
		}
	}
	api.UpdateStatus(map[string]interface{}{
		"desc": form.Desc,
	})
	db.Session.Model(&api).Update(form)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "updated",
		"result":  api,
	})
}

func ApiDeleteApi(ctx iris.Context) {
	var api model.Api
	id, err := ctx.Params().GetUint("id")
	api, err = GetApiItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	db.Session.Delete(&api)
	ctx.StatusCode(iris.StatusNoContent)
}
