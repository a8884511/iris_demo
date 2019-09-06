package admin

import (
	"Goo/app/db"
	"Goo/app/model"
	"Goo/app/util"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

func GetUserItemByID(id uint) (user model.User, err error) {
	result := db.Session.Preload("Group").First(&user, "id = ?", id)
	return user, result.Error
}

func GetUserItemByUsername(username string) (user model.User, err error) {
	result := db.Session.Preload("Group").First(&user, "username = ?", username)
	return user, result.Error
}

func GetUserList(pagination util.Pagination) (users []model.User, count uint, err error) {
	result := db.Session.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Group").Find(&users).Offset(-1).Limit(-1).Count(&count)
	return users, count, result.Error
}

func UserListApi(ctx iris.Context) {
	var users []model.User
	var count uint
	pagination, err := util.GetPagination(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	users, count, err = GetUserList(pagination)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"result": users,
		"count":  count,
	})
}

type UserCreateForm struct {
	Username    string     `json:"username" validate:"required"`
	Password    string     `json:"password" validate:"required"`
	IsActive    bool       `json:"is_active"`
	IsSuperuser bool       `json:"is_superuser"`
	Nickname    string     `json:"nickname"`
	Birthday    *time.Time `json:"birthday"`
	GroupID     uint       `json:"group_id"`
	Desc        string     `json:"desc"`
}

func UserCreateApi(ctx iris.Context) {
	validate := validator.New()
	var form UserCreateForm
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
	var user model.User
	user, err := GetUserItemByUsername(form.Username)
	if err == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "username already exists"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := mapstructure.Decode(form, &user); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := user.SetPassword(form.Password); err != nil {
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
	user.UpdateStatus(map[string]interface{}{
		"desc": form.Desc,
	})
	db.Session.Create(&user)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "created",
		"result":  user,
	})
}

func UserDetailApi(ctx iris.Context) {
	var user model.User
	id, err := ctx.Params().GetUint("id")
	user, err = GetUserItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(user)
}

type UserUpdateForm struct {
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	IsActive    *bool      `json:"is_active"`
	IsSuperuser *bool      `json:"is_superuser"`
	Nickname    string     `json:"nickname"`
	Birthday    *time.Time `json:"birthday"`
	GroupID     uint       `json:"group_id"`
	Desc        string     `json:"desc"`
}

func UserUpdateApi(ctx iris.Context) {
	validate := validator.New()
	var form UserUpdateForm
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
	var user model.User
	id, err := ctx.Params().GetUint("id")
	user, err = GetUserItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if form.Username != "" {
		u, err := GetUserItemByUsername(form.Username)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON(iris.Map{"message": err.Error()})
				return
			}
		} else if u.Username != user.Username {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "username already exists"})
			return
		}
	}
	if form.Password != "" {
		if err := user.SetPassword(form.Password); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"message": err.Error()})
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
	user.UpdateStatus(map[string]interface{}{
		"desc": form.Desc,
	})
	db.Session.Model(&user).Update(form)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "updated",
		"result":  user,
	})
}

func UserDeleteApi(ctx iris.Context) {
	var user model.User
	id, err := ctx.Params().GetUint("id")
	user, err = GetUserItemByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	db.Session.Delete(&user)
	ctx.StatusCode(iris.StatusNoContent)
}
