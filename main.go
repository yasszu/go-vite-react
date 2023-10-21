package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yasszu/go-vite-react/config"
	"github.com/yasszu/go-vite-react/handler"
)

func main() {
	conf := config.NewConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	handler.NewHandler(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         conf.Addr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// TODO: implement Graceful Shutdown
	log.Println("⇨ started on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
