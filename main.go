package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	conf := NewConf()
	h := NewHandler()
	r := NewRouter(h)

	srv := &http.Server{
		Handler:      r,
		Addr:         conf.Addr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("⇨ started on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
