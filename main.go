package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.RequestID,
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		serveFile(w, r)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("⇨ started on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
