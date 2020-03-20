package utils

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/juju/errors"
	"github.com/micro/go-micro/v2/config"
	"time"
)

func InitRedisPool() *redis.Pool {
	host := config.Get("redis", "host").String("192.168.111.151")
	port := config.Get("redis", "port").Int(6379)
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   20,
		IdleTimeout: time.Duration(180) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return nil, errors.Trace(err)
			}
			c.Do("SELECT", 0)
			return c, nil
		},
	}
}
