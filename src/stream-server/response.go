package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
	w.WriteHeader(statusCode)
	io.WriteString(w, errorMsg)
}
