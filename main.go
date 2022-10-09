package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	port = 3000
)

func main() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render(w, r.URL.Path)
	})

	fmt.Printf("→ http server started on http://localhost:%d/\n\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
