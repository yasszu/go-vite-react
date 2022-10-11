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

var ErrFileNotFound = errors.New("file not found")

func OpenFile(fileName string) (fs.File, error) {
	file, err := dist.Open(path.Join(rootPath, fileName))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	stat, _ := file.Stat()
	if stat.IsDir() {
		return nil, ErrFileNotFound
	}

	return file, nil
}
