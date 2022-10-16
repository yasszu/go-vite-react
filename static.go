package main

import (
	"embed"
	"errors"
	"io/fs"
	"path"
)

const (
	rootPath = "dist"
)

//go:embed dist/*
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

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, ErrFileNotFound
	}

	return file, nil
}
