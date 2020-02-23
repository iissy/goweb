package utils

import (
	"database/sql"
	"errors"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"hrefs.cn/src/config"
	"log"
	"os"
)

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", config.String("mysql:hrefs", ""))
	checkErr(err, "sql.Open failed")

	db.SetMaxIdleConns(config.Int("mysql:MaxIdleConns", 5))
	db.SetMaxOpenConns(config.Int("mysql:MaxOpenConns", 50))
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	if config.Bool("mysql:gorp:trace-on", false) {
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
