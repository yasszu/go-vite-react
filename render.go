package main

import (
	"embed"
	"errors"
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

const (
	rootPath = "vite-project/dist"
	rootFile = "index.html"
)

var (
	ErrNotFilePath = errors.New("not file path")
)

func render(w http.ResponseWriter, file string) {
	if err := open(w, file); err != nil {
		log.Println(err)
		if err := open(w, rootFile); err != nil {
			panic(err)
		}
	}
}

//go:embed vite-project/dist/*
var dist embed.FS

func open(w http.ResponseWriter, file string) error {
	f, err := dist.Open(path.Join(rootPath, file))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	stat, _ := f.Stat()
	if stat.IsDir() {
		return ErrNotFilePath
	}

	contentType := "application/octet-stream"
	if mt := mime.TypeByExtension(filepath.Ext(file)); mt != "" {
		contentType = mt
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, f); err != nil {
		return err
	}
	return nil
}
