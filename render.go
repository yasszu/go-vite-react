package main

import (
	"embed"
	"errors"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

const (
	rootPath = "vite-project/dist"
	rootFile = "/index.html"
)

var (
	ErrNotFilePath = errors.New("not file path")
)

func fileServer(w http.ResponseWriter, r *http.Request) {
	render(w, r)
}

func render(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	file, err := open(filePath)
	if err != nil {
		renderRoot(w, r)
		return
	}

	w.Header().Set("Content-Type", contentType(filePath))
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

func renderRoot(w http.ResponseWriter, r *http.Request) {
	file, err := open(rootFile)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

func contentType(filePath string) string {
	ct := "application/octet-stream"
	if mt := mime.TypeByExtension(filepath.Ext(filePath)); mt != "" {
		ct = mt
	}
	return ct
}

//go:embed vite-project/dist/*
var dist embed.FS

func open(fileName string) (fs.File, error) {
	file, err := dist.Open(path.Join(rootPath, fileName))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	stat, _ := file.Stat()
	if stat.IsDir() {
		return nil, ErrNotFilePath
	}

	return file, nil
}
