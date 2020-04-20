package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser handler.
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User function called")
}

// UserLogin handler.
func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uName := p.ByName("user_name")
	io.WriteString(w, uName)
}
