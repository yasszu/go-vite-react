package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yasszu/go-vite-react/render"
)

type Handler struct {
	r render.Render
}

func NewHandler(r *chi.Mux) {
	h := &Handler{
		r: render.NewRender(),
	}
	r.Get("/health", h.Health)
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/hello", h.Hello)
		})
	})
	r.NotFound(h.ServeFile)
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
