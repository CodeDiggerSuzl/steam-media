package session

import (
	"stream-media/src/api/dbops"
	"stream-media/src/api/defs"
	"stream-media/src/api/utils"
	"sync"
	"time"
)

/*
 This file do the following things:
 1. Fetch all session from db when the system starts.
 2. Generate a new session id when a new user logins.
 3. Check the session is expired or not, if expired, return a status code to the front end, and the user need to login again.

Don't to use redis, the use is few and mataining a new module can be tricky.
Use this sync.Map is because the Map is not for concurrent, will panic when more than tow goroutine access the map.
You need to lock before you access the map.
Since the 1.9, Go add a sync Map to solve the map. Read is ok, write will add a global lock(TODO).
*/
var sessionMap *sync.Map // cache

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 { return time.Now().UnixNano() / 1000000 }

// FetchSessionFromDB load session from DB
func FetchSessionFromDB() {
	r, err := dbops.FetchAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		session := v.(*defs.Session)
		sessionMap.Store(k, session)
		return true
	})
}

// GenerateNewSessionID insert into db and generate id
func GenerateNewSessionID(uName string) string {
	id, _ := utils.GenerateUUID()
	currentTime := nowInMilli()
	ttl := currentTime + 30*60*1000 // 30 mins

	session := &defs.Session{UserName: uName, TTL: ttl}

	sessionMap.Store(id, session)
	dbops.InsertSession(id, ttl, uName)
	return id

}

// IsSessExpired judge the session is expired or not
func IsSessExpired(sessionID string) (string, bool) {
	session, ok := sessionMap.Load(sessionID)
	if ok {
		currentTime := nowInMilli()
		if session.(*defs.Session).TTL < currentTime {
			deleteExpiredSession(sessionID)
			return "", true
		}
		return session.(*defs.Session).UserName, false
	}
	return "", true
}

func deleteExpiredSession(sessionID string) {
	sessionMap.Delete(sessionID)
	dbops.DelSessionByID(sessionID)
}
