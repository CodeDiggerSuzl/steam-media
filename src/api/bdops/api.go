package dbops

import (
	"log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// AddUserCredential add a new user.
// The next tow funcs all need the dbConn, the `database/sql` is mainly for long connect sql,not for constant open and close.
// It will waste a lot of resource to create the db.conn.
func AddUserCredential(loginName string, password string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name,password) VALUES(?,?)")
	if err != nil {
		return err
	}

	stmtIns.Exec(loginName, password)
	// close
	defer stmtIns.Close()
	return nil
}

// GetUserCredential get user credential
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT password from users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var password string
	stmtOut.QueryRow(loginName).Scan(&password)
	stmtOut.Close()
	return password, nil
}

//DelUser delete user
func DelUser(loginName string, password string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND password = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	stmtDel.Exec(loginName, password)
	stmtDel.Close()
	return nil
}
