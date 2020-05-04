package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

// NewMiddleWireHandler constructor
func NewMiddleWireHandler(r *httprouter.Router, cc int) http.Handler {
	log.Printf("Creating new middle-wire-handler in stream server module")
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
	// return middleWareHandler{r: r, l: NewConnLimiter(cc)}
}

// Implement of the server method of http
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		log.Printf("Error during Get Connection: %v", m.l.GetConn())
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many request")
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

// RegisterHandlers register handles
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// stream video
	router.GET("/videos/:vid-id", streamingHandler)
	// upload video
	router.POST("/upload/:vid-id", uploadHandler)

	router.GET("/testpage", testPageHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	log.Printf("Listening port 9000, with router :%v", r)
	middleWareHandler := NewMiddleWireHandler(r, 2)
	log.Printf("Listening port 9000, with middleWareHandler :%v", middleWareHandler)
	http.ListenAndServe("localhost:9000", middleWareHandler)
}
