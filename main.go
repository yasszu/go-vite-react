package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yasszu/go-vite-react/pkg/config"
	"github.com/yasszu/go-vite-react/pkg/handler"
	"github.com/yasszu/go-vite-react/pkg/router"
)

func main() {
	c := config.NewConfig()
	h := handler.NewHandler()
	r := router.NewRouter(h)

	srv := &http.Server{
		Handler:      r,
		Addr:         c.Addr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("⇨ started on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
