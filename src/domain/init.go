package domain

import (
	"github.com/go-gorp/gorp"
	"hrefs.cn/src/utils"
)

var dbMap *gorp.DbMap

func InitDb() {
	dbMap = utils.InitDb()
}
