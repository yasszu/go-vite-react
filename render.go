package main

import (
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

func render(w http.ResponseWriter, r *http.Request) {
	file, err := openFile(r.URL.Path)
	if err != nil {
		renderHtml(w, r)
		return
	}

	w.Header().Set("Content-Type", contentType(r.URL.Path))
	w.WriteHeader(http.StatusOK)

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	if _, err := w.Write(bytes); err != nil {
		log.Println(err)
	}
}

func renderHtml(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join(r.URL.Path, "index.html")
	file, err := openFile(filePath)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	if _, err := w.Write(bytes); err != nil {
		log.Println(err)
	}
}

func contentType(filePath string) string {
	ct := "application/octet-stream"
	if mt := mime.TypeByExtension(filepath.Ext(filePath)); mt != "" {
		ct = mt
	}
	return ct
}
