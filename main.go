package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yasszu/go-vite-react/handler"
)

func main() {
	conf := NewConfig()
	h := handler.NewHandler()
	r := NewRouter(h)

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
