package utils

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/juju/errors"
	"hrefs.cn/src/config"
	"time"
)

func InitRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   20,
		IdleTimeout: time.Duration(180) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.String("redis:host", "192.168.111.150"), config.Int("redis:port", 6379)))
			if err != nil {
				return nil, errors.Trace(err)
			}
			c.Do("SELECT", 0)
			return c, nil
		},
	}
}
