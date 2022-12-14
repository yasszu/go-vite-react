package main

import (
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

func RenderFile(w http.ResponseWriter, r *http.Request, fileName string) {
	file, err := OpenFile(fileName)
	if err != nil {
		if err == ErrFileNotFound {
			RenderPage(w, r, fileName)
			return
		}
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", contentType(fileName))
	w.WriteHeader(http.StatusOK)
	write(w, file)
}

func RenderPage(w http.ResponseWriter, r *http.Request, dirName string) {
	filePath := path.Join(dirName, "index.html")
	file, err := OpenFile(filePath)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	write(w, file)
}

func write(w http.ResponseWriter, file io.Reader) {
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	if _, err := w.Write(bytes); err != nil {
		log.Println(err)
	}
}

func contentType(filePath string) string {
	mt := mime.TypeByExtension(filepath.Ext(filePath))
	if mt == "" {
		return "application/octet-stream"
	}
	return mt
}
