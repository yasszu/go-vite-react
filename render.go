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
)

func fileServer(w http.ResponseWriter, r *http.Request) {
	render(w, r)
}

func render(w http.ResponseWriter, r *http.Request) {
	file, err := open(r.URL.Path)
	if err != nil {
		renderHtml(w, r)
		return
	}

	w.Header().Set("Content-Type", contentType(r.URL.Path))
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, file)
}

func renderHtml(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join(r.URL.Path, "index.html")
	file, err := open(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, file)
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

var ErrNotFilePath = errors.New("not file path")

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
