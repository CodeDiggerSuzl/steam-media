package dbops

import (
	"database/sql"
	"log"
	"stream-media/src/api/defs"
	"stream-media/src/api/utils"
	"time"

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

// video related apis

// AddNewVideo add new video_info into db
func AddNewVideo(authorID int, name string) (*defs.VideoInfo, error) {
	// create UUID
	vID, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	// Can't change the exact time
	createTime := t.Format("Jan 02 2006, 15:04:05")
	stmtInsert, err := dbConn.Prepare("INSERT INTO video_info (id,author_id,name,display_ctime)VALUES(?,?,?,?)")
	defer stmtInsert.Close()
	if err != nil {
		return nil, err
	}
	_, err = stmtInsert.Exec(vID, authorID, name, createTime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{ID: vID, AuthorID: authorID, Name: name, DisplayCreateTime: createTime}
	return res, nil
}

// GetVideoInfo get video info from db
func GetVideoInfo(vID string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")

	defer stmtOut.Close()

	var (
		authorID int
		dct      string
		name     string
	)
	// err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	err = stmtOut.QueryRow(vID).Scan(&authorID, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.VideoInfo{ID: vID, AuthorID: authorID, Name: name, DisplayCreateTime: dct}
	// res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

// DelVideoInfo delete video info by id
func DelVideoInfo(vID string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	defer stmtDel.Close()

	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vID)
	if err != nil {
		return err
	}
	return nil
}
