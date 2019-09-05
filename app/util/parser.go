package util

import (
	"github.com/juju/errors"
	"github.com/kataras/iris"
)

type Pagination struct {
	Offset uint
	Limit  uint
}

func GetPagination(ctx iris.Context) (pagination Pagination, err error) {
	offset := ctx.URLParamIntDefault("offset", 0)
	limit := ctx.URLParamIntDefault("limit", 10)
	if offset < 0 || limit < 0 {
		err = errors.New("Query params 'offset' or 'limit' must greater than zero")
		return
	}
	pagination = Pagination{
		Offset: uint(offset),
		Limit:  uint(limit),
	}
	return
}
