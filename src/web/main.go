package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}

// RegisterHandler register handlers
func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	// for log out action and stuff
	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)

	router.POST("/api", apiHandler)

	router.POST("/upload/:vid-id", proxyHandler)

	// warper of http.files
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}
