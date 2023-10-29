package render

import (
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
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

func (r *Render) TryRenderFile(w http.ResponseWriter, req *http.Request, fileName string) {
	file, err := r.dist.OpenFile(fileName)
	if err != nil {
		r.renderPage(w, req, "/index.html")
		return
	}
	r.renderFile(w, file, fileName)
}

func (r *Render) renderPage(w http.ResponseWriter, req *http.Request, filePath string) {
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

func (r *Render) renderFile(w http.ResponseWriter, file fs.File, fileName string) {
	w.Header().Set("Content-Type", contentType(fileName))
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
