package main

import (
	"net/http"
	"stream-media/src/scheduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandler register handler
func RegisterHandler() *httprouter.Router {

	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	// use goroutine
	// c := make(chan int) make the func blocking mode
	go taskrunner.Start()
	r := RegisterHandler()
	// <-c
	http.ListenAndServe(":9001", r)
}
