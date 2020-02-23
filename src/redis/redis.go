package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/juju/errors"
	"hrefs.cn/src/utils"
)

var pool *redis.Pool

func InitDb() {
	pool = utils.InitRedisPool()
}

func Get(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", errors.Trace(err)
	}

	return result, nil
}

func Set(key string, value string) error {
	conn := pool.Get()
	defer conn.Close()
	err := conn.Send("SET", key, value)
	if err != nil {
		return errors.Trace(err)
	}

	expireTime := 60 * 60 * 6
	err = conn.Send("EXPIRE", key, expireTime)
	if err != nil {
		return errors.Trace(err)
	}

	err = conn.Flush()
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}
