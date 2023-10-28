package web

import (
	"embed"
	"errors"
	"io/fs"
	"path"
)

const (
	rootPath = "vite-project/dist"
)

var (
	//go:embed vite-project/dist/*
	dist embed.FS

	ErrFileNotFound = errors.New("file not found")
)

type Dist struct {
	dist embed.FS
}

func NewDist() Dist {
	return Dist{dist: dist}
}

func (d *Dist) OpenFile(fileName string) (fs.File, error) {
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
