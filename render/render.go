package render

import (
	"errors"
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"

	"github.com/yasszu/go-vite-react/web"
)

type Render struct {
	dist web.Dist
}

func NewRender() Render {
	return Render{
		dist: web.NewDist(),
	}
}

func (r *Render) RenderFile(w http.ResponseWriter, req *http.Request, fileName string) {
	file, err := r.dist.OpenFile(fileName)
	if err != nil {
		if errors.Is(err, web.ErrFileNotFound) {
			r.renderPage(w, req, fileName)
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

func (r *Render) renderPage(w http.ResponseWriter, req *http.Request, dirName string) {
	filePath := path.Join(dirName, "index.html")
	file, err := r.dist.OpenFile(filePath)
	if err != nil {
		log.Println(err)
		http.NotFound(w, req)
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
