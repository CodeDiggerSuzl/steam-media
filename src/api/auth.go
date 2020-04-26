package main

import (
	"log"
	"net/http"
	"stream-media/src/api/defs"
	"stream-media/src/api/session"
)

const (
	// In http, if a header starts with 'X' is customize header
	HEADER_FILED_SESSION = "X-Session-ID"
	HEADER_FILED_UNAME   = "X-Session-Name"
)

// check the session is validate or not
func validateUserSession(r *http.Request) bool {
	log.Printf("validUserSession: %v", r)

	sessionID := r.Header.Get(HEADER_FILED_SESSION)
	// check the customize header
	if len(sessionID) == 0 {
		return false
	}
	// is expired or not
	userName, ok := session.IsSessExpired(sessionID)
	if ok {
		return false
	}
	log.Printf("validUserSession userName: %v", userName)
	r.Header.Add(HEADER_FILED_UNAME, userName)
	return true
}

// ValidateUser valid user
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	userName := r.Header.Get(HEADER_FILED_UNAME)
	log.Printf("ValidateUser userName: %v", userName)
	if len(userName) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
