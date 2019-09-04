package middleware

import (
	"github.com/casbin/casbin"
	"github.com/kataras/iris"
)

var ACL *casbin.Enforcer

var ACLInitialize = func() {
	ACL = casbin.NewEnforcer("config/acl/model.conf", "config/acl/policy.csv")
}

var ACLMiddleware = func(ctx iris.Context) {
	ctx.Next()
}
