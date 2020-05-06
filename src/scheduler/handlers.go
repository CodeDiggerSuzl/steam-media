package main

import (
	"net/http"
	"stream-media/src/scheduler/dbops"

	"github.com/julienschmidt/httprouter"
)

func vidDelRecHandel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResp(w, 400, "video id should not be empty")
		return
	}
	err := dbops.AddVideoDeleteRecord(vid)
	if err != nil {
		sendResp(w, 500, "Internal server error")
		return
	}
	sendResp(w, 200, "")
	return
}
