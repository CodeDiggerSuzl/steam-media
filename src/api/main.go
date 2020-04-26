package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

// NewMiddleWareHandler factory
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

// Implements of ServeHTTP method, so the middleware handler can response intercept the reqest.
// Each request need to check session and do auth things
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	log.Printf("ServeHTTP: %v", r)
	validateUserSession(r) // TODO this func returns a bool
	m.r.ServeHTTP(w, r)
}

// main func file only put some defs, logic code should put in other files.
func main() {
	r := RegisterHandlers()
	// intercept each request
	middleWareHandler := NewMiddleWareHandler(r)
	_ = http.ListenAndServe(":8000", middleWareHandler)
}

// RegisterHandlers router.
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	log.Printf("RegisterHandlers: %v", router)
	// Create user handler,use closure.
	router.POST("/user", CreateUser)

	// 	 With param
	router.POST("/user/:user_name", UserLogin)
	return router
}
