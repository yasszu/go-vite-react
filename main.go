package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	port := 3000
	fmt.Printf("→ http server started on http://localhost:%d/\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
