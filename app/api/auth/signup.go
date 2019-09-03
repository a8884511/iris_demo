package auth

import (
	"Goo/app/db"
	"Goo/app/model"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type SignUpForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func SignUpApi(ctx iris.Context) {
	validate := validator.New()
	var signUpForm SignUpForm
	if err := ctx.ReadJSON(&signUpForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := validate.Struct(signUpForm); err != nil {
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
	if result := db.Session.First(&user, "username = ?", signUpForm.Username); result.Error == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Username already exists"})
		return
	}
	user = model.User{Username: signUpForm.Username}
	user.SetPassword(signUpForm.Password)
	db.Session.Create(&user)
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Sign Up"})
}
