package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, nil)
	} else {
		f, h, _ := r.FormFile("image")

		filename := h.Filename
		defer f.Close()

		t, _ := os.Create(UPLOAD_DIR + "/" + filename)
		defer t.Close()

		_, err := io.Copy(t, f)
		if err != nil {

		}
	}
}
