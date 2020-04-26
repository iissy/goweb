package domain

import (
	"github.com/go-gorp/gorp"
	"github.com/iissy/goweb/src/utils"
)

var dbMap *gorp.DbMap

func InitDb() {
	dbMap = utils.InitDb()
}
