package auth

import (
	"Goo/app/db"
	"Goo/app/model"
	"Goo/app/util"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type SignInForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func SignInApi(ctx iris.Context) {
	validate := validator.New()
	var form SignInForm
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
	if result := db.Session.First(&user, "username = ?", form.Username); result.Error != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "username does not exist"})
		return
	}
	if err := user.CheckPassword(form.Password); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "password is incorrect"})
		return
	}
	tokenStr, err := util.GenerateJWToken(user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "generate token error"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "sign in", "token": tokenStr})
}
