package dbops

import (
	"log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// func openConn() *sql.DB {
// 	dbConn, err := sql.Open("mysql", "root:mysql@suz1@localhost:3306/stream_video?charset=utf8")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return dbConn
// }

// AddUserCredential add a new user.
// The next tow funcs all need the dbConn, the `database/sql` is mainly for long connect sql,not for constant open and close.
// It will waste a lot of resource to create the db.conn.
func AddUserCredential(loginName string, passward string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO user (login_name,passward) values(?,?)")
	if err != nil {
		return err
	}

	stmtIns.Exec(loginName, passward)
	// close
	defer stmtIns.Close()
	return nil
}

// GetUserCredential get user credential
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT passward from user where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var passward string
	stmtOut.QueryRow(loginName).Scan(&passward)
	stmtOut.Close()
	return passward, nil
}

//DelUser delete user
func DelUser(loginName string, passward string) error {
	stmtDel, err := dbConn.Prepare("DELETE FORM user where login_name = ? and passward = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	stmtDel.Exec(loginName, passward)
	stmtDel.Close()
	return nil
}
