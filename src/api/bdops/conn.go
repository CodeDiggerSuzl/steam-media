package dbops

import (
	"database/sql"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// *The global variable in the current package.
var (
	dbConn *sql.DB
	err    error
)

// 	defer dbConn.Close()
// ! init function is a special func
// * https://zhuanlan.zhihu.com/p/34211611

func init() {
	// ! can't use :=
	dbConn, err = sql.Open("mysql", "root:mysql@suz1@tcp(127.0.0.1:3306)/stream_video_server")
	if err != nil {
		panic(err.Error())
	}
}
