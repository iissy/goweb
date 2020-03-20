package utils

import (
	"database/sql"
	"errors"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/config"
	"log"
	"os"
)

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", config.Get("mysql", "hrefs").String(""))
	checkErr(err, "sql.Open failed")

	db.SetMaxIdleConns(config.Get("mysql", "MaxIdleConns").Int(5))
	db.SetMaxOpenConns(config.Get("mysql", "MaxOpenConns").Int(50))
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	if config.Get("mysql", "gorp", "trace-on").Bool(false) {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "[SQL]:", log.Lmicroseconds))
	}

	return dbMap
}

func BuildSqlArgs(args ...interface{}) ([]interface{}, error) {
	newArgs := make([]interface{}, 0)
	addEleFun := func(ele interface{}) {
		newArgs = append(newArgs, ele)
		return
	}
	for _, arg := range args {
		switch v := arg.(type) {
		case string, int, int32, int64, bool, *string, *int, *int32, *int64:
			addEleFun(v)
		case []string:
			for _, e := range v {
				addEleFun(e)
			}
		case []int:
			for _, e := range v {
				addEleFun(e)
			}
		case []int32:
			for _, e := range v {
				addEleFun(e)
			}
		case []int64:
			for _, e := range v {
				addEleFun(e)
			}
		case []*string:
			for _, e := range v {
				addEleFun(e)
			}
		default:
			return nil, errors.New("miss type")
		}
	}
	return newArgs, nil
}
