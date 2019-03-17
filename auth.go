package main

import (
	"net/http"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"iissy.com/utils"
)

func basicAuth(h context.Handler) context.Handler {
	return func(ctx iris.Context) {
		if ok := utils.Check(ctx); ok {
			h(ctx)
		} else {
			http.Error(ctx.ResponseWriter(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
