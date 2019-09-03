package auth

import (
	"Goo/app/util"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type TokenForm struct {
	Token string `json:"token" validate:"required"`
}

func TokenCheckApi(ctx iris.Context) {
	validate := validator.New()
	var tokenForm TokenForm
	if err := ctx.ReadJSON(&tokenForm); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	if err := validate.Struct(tokenForm); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"message": err.Error()})
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	claim, err := util.ParseJWToken(tokenForm.Token)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.JSON(claim)
}
