package auth

import (
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func LoginApi(ctx iris.Context) {
	validate := validator.New()

	var loginForm LoginForm
	if err := ctx.ReadJSON(&loginForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	if err := validate.Struct(loginForm); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}

		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(iris.Map{"message": "Validated successfully"})
}
