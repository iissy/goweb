package utils

import (
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

func checkErr(err error, msg string) {
	if err != nil {
		logrus.Fatalln(msg, err)
	}
}

func WriteErrorLog(path string, err error) bool {
	if err != nil {
		logrus.Errorf("%s, url = %s", errors.ErrorStack(err), path)
		return true
	} else {
		return false
	}
}
