package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yasszu/go-vite-react/pkg"
)

func main() {
	conf := pkg.NewConfig()
	h := pkg.NewHandler()
	r := pkg.NewRouter(h)

	srv := &http.Server{
		Handler:      r,
		Addr:         conf.Addr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("⇨ started on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
