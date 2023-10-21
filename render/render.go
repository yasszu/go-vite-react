package render

import (
	"embed"
	"errors"
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

const (
	rootPath = "dist"
)

var (
	//go:embed dist/*
	dist embed.FS

	ErrFileNotFound = errors.New("file not found")
)

type Render struct {
	dist embed.FS
}

func NewRender() Render {
	return Render{
		dist: dist,
	}
}

func (r *Render) RenderFile(w http.ResponseWriter, req *http.Request, fileName string) {
	file, err := openFile(fileName)
	if err != nil {
		if errors.Is(err, ErrFileNotFound) {
			renderPage(w, req, fileName)
			return
		}
		log.Println(err)
		http.NotFound(w, req)
		return
	}

	w.Header().Set("Content-Type", contentType(fileName))
	w.WriteHeader(http.StatusOK)
	write(w, file)
}

func renderPage(w http.ResponseWriter, req *http.Request, dirName string) {
	filePath := path.Join(dirName, "index.html")
	file, err := openFile(filePath)
	if err != nil {
		log.Println(err)
		http.NotFound(w, req)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	write(w, file)
}

func openFile(fileName string) (fs.File, error) {
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
