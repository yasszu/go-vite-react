package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	conf := NewConfig()
	h := NewHandler()
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
