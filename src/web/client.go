package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error
	var req *http.Request

	switch b.Method {
	case http.MethodGet, http.MethodDelete:
		req, _ = http.NewRequest(b.Method, b.Url, nil)
	case http.MethodPost:
		req, _ = http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
	}

	req.Header = r.Header
	resp, err = httpClient.Do(req)
	if err != nil {
		log.Printf("Error during do GET mothod: %v", err)
		return
	}
	normalResponse(w, resp)
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error during io.ReadAll: %v", err)
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}
	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}
