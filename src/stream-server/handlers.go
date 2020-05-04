package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func streamingHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	videoPath := VIDEO_DIR + vid

	video, err := os.Open(videoPath)
	defer video.Close()
	if err != nil {
		log.Printf("Error during open file : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Error during opening file")
		return
	}

	w.Header().Set("Content_Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
}
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("Error during upload handler: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, "File is too large, should less than 100M")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error during parse file %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error during ioutil.ReadAll %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
	}
	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0755)
	if err != nil {
		log.Printf("Error during write file %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload successfully")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./video/upload.html")
	t.Execute(w, nil)
}
