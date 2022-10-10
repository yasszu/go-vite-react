package main

import (
	"embed"
	"errors"
	"io/fs"
	"path"
)

const (
	rootPath = "vite-project/dist"
)

//go:embed vite-project/dist/*
var dist embed.FS

var ErrNotFoundFile = errors.New("not found file")

func openFile(fileName string) (fs.File, error) {
	file, err := dist.Open(path.Join(rootPath, fileName))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	stat, _ := file.Stat()
	if stat.IsDir() {
		return nil, ErrNotFoundFile
	}

	return file, nil
}
