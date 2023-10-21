package handler

import (
	"net/http"

	"github.com/yasszu/go-vite-react/render"
)

type Handler struct {
	r render.Render
}

func NewHandler() *Handler {
	return &Handler{
		r: render.NewRender(),
	}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func (h *Handler) Hello(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World!"))
}

func (h *Handler) ServeFile(w http.ResponseWriter, r *http.Request) {
	h.r.RenderFile(w, r, r.URL.Path)
}
