package dbops

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "mysql@suz1"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "server"
)

func TestDBConn(t *testing.T) {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Open failed")
		t.Errorf("err %v", err)
		return
	}
}
