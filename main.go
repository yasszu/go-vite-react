package main

import (
	"fmt"
	"net/http"

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
	r.NotFound(fileServer)

	port := 3000
	fmt.Printf("\n⇨ http://localhost:%d/\n\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
