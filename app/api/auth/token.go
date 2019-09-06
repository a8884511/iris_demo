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
	var form TokenForm
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
	claim, err := util.ParseJWToken(form.Token)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(claim)
}
