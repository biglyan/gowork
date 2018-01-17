package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// UploadDir...
const (
	UploadDir = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, nil)
	} else {
		f, h, _ := r.FormFile("image")

		filename := h.Filename
		defer f.Close()

		t, _ := os.Create(UploadDir + "/" + filename)
		defer t.Close()

		_, err := io.Copy(t, f)
		if err != nil {
			log.Fatal(err)
			return
		}

		http.Redirect(w, r, "view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := UploadDir + "/" + imageID
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UploadDir)
	if err != nil {
		log.Fatal("read dir error ", err)
		return
	}

	locals := make(map[string]interface{})
	images := []string{}

	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}

	locals["images"] = images

	t, err := template.ParseFiles("list.html")
	if err != nil {
		log.Fatal("template error", err.Error())
		return
	}

	t.Execute(w, locals)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/", listHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
