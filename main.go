package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	conf := NewConf()
	h := NewHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Get("/health", h.Health)
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/hello", h.Hello)
		})
	})
	r.NotFound(h.ServeFile)

	srv := &http.Server{
		Handler:      r,
		Addr:         conf.Addr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("⇨ started on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
