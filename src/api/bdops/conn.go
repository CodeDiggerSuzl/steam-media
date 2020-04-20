package dbops

import (
	"database/sql"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// The global variable in the package
var (
	dbConn *sql.DB
	err    error
)

// DbWorker for db
type DbWorker struct {
	// mysql data source name
	DSN string
}

func initDBConn() {
	dbWorker := DbWorker{DSN: "root:mysql@suz1@tcp(127.0.0.1:3306)/stream_video_server"}
	dbConn, err = sql.Open("mysql", dbWorker.DSN)
	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()
}
