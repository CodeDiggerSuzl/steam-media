package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// main func file only put some defs, logic code should put in other files.
func main() {

	db, err := sql.Open("mysql", "root:mysql@suz1@tcp(127.0.0.1:3306)/stream_video_server")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Open failed")

		return
	}
	fmt.Println(db)
}

// RegisterHandlers router.
// func RegisterHandlers() *httprouter.Router {
// 	router := httprouter.New()

// 	// Create user handler,use colsure.
// 	router.POST("/user", CreateUser)

// 	// With param
// 	router.POST("/user/:user_name", UserLogin)
// 	return router
// }
