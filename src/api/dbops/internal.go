package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"stream-media/src/api/defs"
	"sync"
)

// InsertSession Write session to DB
func InsertSession(sessionID string, ttl int64, userName string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?,?,?)")
	defer stmtIns.Close()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sessionID, ttlStr, userName)
	if err != nil {
		log.Printf("Err during inserting into session %v", err)
		return err
	}
	return nil
}

// FetchSessionByID  Load session by id
func FetchSessionByID(sessionID string) (*defs.Session, error) {
	session := &defs.Session{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	defer stmtOut.Close()

	if err != nil {
		return nil, err
	}
	var ttl, userName string
	stmtOut.QueryRow(sessionID).Scan(&ttl, &userName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		session.TTL = res
		session.UserName = userName
	} else {
		return nil, err
	}
	return session, nil
}

// FetchAllSessions load all session from DB
func FetchAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	defer stmtOut.Close()

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id, ttlStr, loginName string
		if err := rows.Scan(&id, &ttlStr, &loginName); err != nil {
			log.Printf("Fetch sessions error: %s", err)
			break
		}
		if ttl, er := strconv.ParseInt(ttlStr, 10, 64); er != nil {
			session := &defs.Session{UserName: loginName, TTL: ttl}
			m.Store(id, session)
			log.Printf("session id %s, ttl: %d", id, session.TTL)
		}
	}
	return m, nil
}

// DelSessionByID delete by session id
func DelSessionByID(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	defer stmtOut.Close()
	if err != nil {
		return err
	}
	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}
	return nil
}
