package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

// NewMiddleWireHandler constructor
func NewMiddleWireHandler(r *httprouter.Router, cc int) http.Handler {
	return middleWareHandler{r: r, l: newConnLimiter(cc)}
}

// Implement of the server method of http
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many request")
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}

// RegisterHandlers register handless
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// stream video
	router.GET("/videos/:vid-id", streamingHandler)
	// upload video
	router.POST("/upload/:vid-id", uploadHandler)

	return router
}
