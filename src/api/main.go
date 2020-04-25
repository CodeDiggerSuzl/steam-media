package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// main func file only put some defs, logic code should put in other files.
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

// RegisterHandlers router.
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// Create user handler,use colsure.
	router.POST("/user", CreateUser)

	// 	 With param
	router.POST("/user/:user_name", UserLogin)
	return router
}
