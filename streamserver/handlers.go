package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid + ".mp4"

	video, err := os.Open(vl)
	if err != nil {
		senErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		senErrorResponse(w, http.StatusBadRequest, "Upload Failed,File is too Large.")
		return
	}
	//参数key对应form表单的name
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to open file : %v", err)
		senErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Printf("Read file error: %v", err)
		senErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	fn := p.ByName("vid-id") + ".mp4"

	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)

	if err != nil {
		log.Printf("Write file error: %v", err)
		senErrorResponse(w, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded Successfully.")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./server/templates/upload.html")
	t.Execute(w, nil)
}
