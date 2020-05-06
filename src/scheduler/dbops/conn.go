package dbops

import (
	"database/sql"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:mysql@suz1@tcp(127.0.0.1:3306)/stream_video_server")
	if err != nil {
		panic(err.Error())
	}
}
