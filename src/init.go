package src

import (
	"hrefs.cn/src/config"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/redis"
)

func init() {
	config.LoadConfigs()
	domain.InitDb()
	redis.InitDb()
}
