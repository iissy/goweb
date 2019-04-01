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
			// name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
			// if name == "main.articlelist" {
			// 	log.Println(name)
			// 	h(ctx)
			// } else {
			// 	log.Println("jk")
			// 	result := struct{ Success bool }{false}
			// 	ctx.JSON(result)
			// }
			h(ctx)
		} else {
			http.Error(ctx.ResponseWriter(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
