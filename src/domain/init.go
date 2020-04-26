package domain

import (
	"github.com/go-gorp/gorp/v3"
	"github.com/iissy/goweb/src/utils"
)

var dbMap *gorp.DbMap

func InitDb() {
	dbMap = utils.InitDb()
}
