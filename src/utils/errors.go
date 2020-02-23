package utils

import (
	"github.com/juju/errors"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"log"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func WriteErrorLog(ctx iris.Context, err error) bool {
	if err != nil {
		golog.Errorf("%s, url = %s", errors.ErrorStack(err), ctx.Path())
		return true
	} else {
		return false
	}
}
