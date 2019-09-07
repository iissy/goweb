package access

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"iissy.com/src/utils"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", utils.SQLDB)
	if err != nil {
		panic(err)
	}

	return
}
