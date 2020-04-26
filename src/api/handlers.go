package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"stream-media/src/api/dbops"
	"stream-media/src/api/defs"
	"stream-media/src/api/session"

	"github.com/julienschmidt/httprouter"
)

// CreateUser handler.
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("CreateUser: %v", res)
	uBody := &defs.UserCredential{}
	// check err
	if err := json.Unmarshal(res, uBody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(uBody.UserName, uBody.Password); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	// Add a new session
	id := session.GenerateNewSessionID(uBody.UserName)
	signUp := &defs.SignedUp{Success: true, SessionID: id}
	resp, err := json.Marshal(signUp)
	if err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	sendNormalResponse(w, string(resp), 201)
}

// UserLogin handler.
func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uName := p.ByName("user_name")
	io.WriteString(w, uName)
}
