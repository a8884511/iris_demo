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

func GetRoleItemByID(id uint) (role model.Role, err error) {
	result := db.Session.Preload("Group").Preload("Apis").First(&role, "id = ?", id)
	return role, result.Error
}

func GetRoleItemByName(name string) (role model.Role, err error) {
	result := db.Session.Preload("Group").Preload("Apis").First(&role, "name = ?", name)
	return role, result.Error
}

func GetRoleList(pagination util.Pagination) (roles []model.Role, count uint, err error) {
	result := db.Session.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Group").Preload("Apis").Find(&roles).Offset(-1).Limit(-1).Count(&count)
	return roles, count, result.Error
}

func GetApiListByIDS(ids []uint) (apis []*model.Api, err error) {
	result := db.Session.Find(&apis, "id IN (?)", ids)
	return apis, result.Error
}

func RoleListApi(ctx iris.Context) {
	var roles []model.Role
	var count uint
	pagination, err := util.GetPagination(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	roles, count, err = GetRoleList(pagination)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"result": roles,
		"count":  count,
	})
}

type RoleCreateForm struct {
	Name    string `json:"name" validate:"required"`
	GroupID uint   `json:"group_id"`
	ApiIDS  []uint `json:"api_ids"`
	Desc    string `json:"desc"`
}

func RoleCreateApi(ctx iris.Context) {
	validate := validator.New()
	var form RoleCreateForm
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
	var role model.Role
	role, err := GetRoleItemByName(form.Name)
	if err == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "role name already exists"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := mapstructure.Decode(form, &role); err != nil {
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
	if form.ApiIDS != nil {
		apis, err := GetApiListByIDS(form.ApiIDS)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": err.Error()})
			return
		}
		role.Apis = apis
	}
	role.UpdateStatus(map[string]interface{}{
		"desc": form.Desc,
	})
	db.Session.Create(&role)
	//db.Session.Exec("DELETE FROM casbin_rule WHERE v0 = ?", role.ID)
	for _, api := range role.Apis {
		db.Session.Exec("INSERT INTO casbin_rule ('p_type', 'v0', 'v1', 'v2', 'v3') VALUES (?, ?, ?, ?, ?)", "p", role.ID, role.GroupID, api.Url, api.Method)
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "created",
		"result":  role,
	})
}

func RoleDetailApi(ctx iris.Context) {
	var role model.Role
	id, err := ctx.Params().GetUint("id")
	role, err = GetRoleItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(role)
}

type RoleUpdateForm struct {
	Name    string `json:"name"`
	GroupID uint   `json:"group_id"`
	ApiIDS  []uint `json:"api_ids"`
	Desc    string `json:"desc"`
}

func RoleUpdateApi(ctx iris.Context) {
	validate := validator.New()
	var form RoleUpdateForm
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
	var role model.Role
	id, err := ctx.Params().GetUint("id")
	role, err = GetRoleItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if form.Name != "" {
		r, err := GetRoleItemByName(form.Name)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON(iris.Map{"message": err.Error()})
				return
			}
		} else if r.Name != role.Name {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "role name already exists"})
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
	if form.ApiIDS != nil {
		apis, err := GetApiListByIDS(form.ApiIDS)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": err.Error()})
			return
		}
		role.Apis = apis
	}
	role.UpdateStatus(map[string]interface{}{
		"desc": form.Desc,
	})
	db.Session.Model(&role).Update(form)
	db.Session.Exec("DELETE FROM casbin_rule WHERE v0 = ?", role.ID)
	for _, api := range role.Apis {
		db.Session.Exec("INSERT INTO casbin_rule ('p_type', 'v0', 'v1', 'v2', 'v3') VALUES (?, ?, ?, ?, ?)", "p", role.ID, role.GroupID, api.Url, api.Method)
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "updated",
		"result":  role,
	})
}

func RoleDeleteApi(ctx iris.Context) {
	var role model.Role
	id, err := ctx.Params().GetUint("id")
	role, err = GetRoleItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	db.Session.Delete(&role)
	ctx.StatusCode(iris.StatusNoContent)
}
