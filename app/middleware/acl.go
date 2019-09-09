package middleware

import (
	"Goo/app/plugin"
	"Goo/app/util"
	"github.com/juju/errors"
	"github.com/kataras/iris"
)

var CasbinMiddleware = func(ctx iris.Context) {
	user, err := util.GetCurrentUser(ctx)
	if err != nil {
		CasbinErrorHandler(ctx, err)
		ctx.StopExecution()
		return
	}
	if *(user.IsSuperuser) != true {
		ok, err := plugin.Enforcer.Enforce(user.ID, user.GroupID, ctx.Path(), ctx.Method())
		if err != nil {
			CasbinErrorHandler(ctx, err)
			ctx.StopExecution()
			return
		}
		if !ok {
			err := errors.New("No Permission")
			CasbinErrorHandler(ctx, err)
			ctx.StopExecution()
			return
		}
	}
	ctx.Next()
}

func CasbinErrorHandler(ctx iris.Context, err error) {
	ctx.StatusCode(iris.StatusForbidden)
	ctx.JSON(iris.Map{"message": err.Error()})
}
