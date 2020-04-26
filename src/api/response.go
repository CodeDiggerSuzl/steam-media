package main

import (
	"encoding/json"
	"io"
	"net/http"
	"stream-media/src/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HTTPSC)

	respStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(respStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
