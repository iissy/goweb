package middleware

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"time"
)

func Trace(ctx iris.Context) {
	defer trace(ctx.Path())()
	ctx.Next()
}

func trace(path string) func() {
	start := time.Now()
	return func() {
		golog.Infof("%s %s", time.Since(start), path)
	}
}
