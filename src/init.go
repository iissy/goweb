package src

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/redis"
	"log"
)

const (
	defaultConfigPath = "conf/config.json"
)

func init() {
	loadConfig()
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
