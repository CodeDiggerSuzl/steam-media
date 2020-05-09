package main

import (
	"net/http"
	"stream-media/src/scheduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandler register handler
func RegisterHandler() *httprouter.Router {

	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandel)
	return router
}

func main() {
	go taskrunner.Start()
	r := RegisterHandler()
	http.ListenAndServe(":9001", r)
}
