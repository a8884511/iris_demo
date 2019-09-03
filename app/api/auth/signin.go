package auth

import (
	"Goo/app/db"
	"Goo/app/model"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type SignInForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//TODO 封装校验函数
func SignInApi(ctx iris.Context) {
	validate := validator.New()
	var signInForm SignInForm
	if err := ctx.ReadJSON(&signInForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := validate.Struct(signInForm); err != nil {
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
	if result := db.Session.First(&user, "username = ?", signInForm.Username); result.Error != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Username is incorrect."})
		return
	}

	if !user.CheckPassword(signInForm.Password) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Password is incorrect."})
		return
	}

	ctx.JSON(iris.Map{"message": "Sign In"})
}
