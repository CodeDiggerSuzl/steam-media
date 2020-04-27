package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}

// RegisterHandlers register handless
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamingHandler)
	router.POST("/upload/:vid-id", uploadHanler)
	return router
}
