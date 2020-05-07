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
	router.POST("/", homeHandler)

	router.GET("/homepage", userHomeHandler)
	router.POST("/homepage", userHomeHandler)

	router.POST("/api", apiHandler)

	router.POST("/upload/:vid-id", proxyHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("../templates"))

	return router
}
