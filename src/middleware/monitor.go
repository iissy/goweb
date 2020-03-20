package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"time"
)

func TraceWeb(ctx iris.Context) {
	defer trace(ctx.Path())()
	ctx.Next()
}

func TraceApi(ctx *gin.Context) {
	defer trace(ctx.Request.URL.Path)()
	ctx.Next()
}

func trace(path string) func() {
	start := time.Now()
	return func() {
		golog.Infof("%s %s", time.Since(start), path)
	}
}
