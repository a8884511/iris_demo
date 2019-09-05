package plugin

import (
	"Goo/app/db"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
)

var Enforcer *casbin.Enforcer

func CasbinInitialize() {
	adapter, err := gormadapter.NewAdapterByDB(db.Session)
	if err != nil {
		panic(err.Error())
	}
	Enforcer, err = casbin.NewEnforcer("config/acl/model.conf", adapter)
	if err != nil {
		panic(err.Error())
	}
}
