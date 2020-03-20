package utils

import (
	"github.com/juju/errors"
	"github.com/kataras/golog"
	"log"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func WriteErrorLog(path string, err error) bool {
	if err != nil {
		golog.Errorf("%s, url = %s", errors.ErrorStack(err), path)
		return true
	} else {
		return false
	}
}
