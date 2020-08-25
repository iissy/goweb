package src

import (
	"github.com/iissy/goweb/src/domain"
	"github.com/iissy/goweb/src/redis"
	"github.com/iissy/goweb/src/utils"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"log"
)

const (
	defaultConfigPath = "conf/config.json"
)

func init() {
	loadConfig()
	utils.InitLog()
	domain.InitDb()
	redis.InitDb()
}

func loadConfig() {
	if err := config.Load(file.NewSource(
		file.WithPath(defaultConfigPath),
	)); err != nil {
		log.Panic(err)
	}
}
