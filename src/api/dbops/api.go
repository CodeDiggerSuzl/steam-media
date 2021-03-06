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
		log.Printf("Error during add user credential: %v", err)
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
		log.Printf("GetUserCredential occurs error: %v", err)
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
		log.Printf("DelUser occurs error: %v", err)
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
	// Can't change the exact time TODO exact time
	createTime := t.Format("Jan 02 2006, 15:04:05")
	stmtInsert, err := dbConn.Prepare("INSERT INTO video_info (id,author_id,name,display_ctime)VALUES(?,?,?,?)")
	defer stmtInsert.Close()
	if err != nil {
		log.Printf("AddNewVideo occurs error: %v", err)
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
		log.Printf("GetVideoInfo: err during scan video info %v ", err)
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
		log.Printf("Error during DelVideoInfo %v", err)
		return err
	}
	return nil
}

/*comments api*/

// AddNewComments add new comments on video
func AddNewComments(vID string, authorID int, content string) error {
	commentID, err := utils.GenerateUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id,video_id,author_id,content) VALUES (?,?,?,?)")
	defer stmtIns.Close()
	_, err = stmtIns.Exec(commentID, vID, authorID, content)
	if err != nil {
		log.Printf("AddNewComments error happens: %v", err)
		return err
	}
	return nil
}

// ListComments list all comments of a video
func ListComments(vID string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(
		`SELECT comments.id, users.login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	defer stmtOut.Close()

	var res []*defs.Comment
	rows, err := stmtOut.Query(vID, from, to)

	if err != nil {
		log.Printf("ListComments error %v", err)
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			log.Printf("ListComments error rows.Next():  %v", err)
			return res, err
		}
		c := &defs.Comment{ID: id, VideoID: vID, Author: name, Content: content}
		res = append(res, c)
	}
	return res, nil
}
