package main

import (
	"log"
	"net/http"
	"time"

	"github.com/c0rrzin/router"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
	DefRoutes()
	router.RouteAll()
	log.Fatal(s.ListenAndServe())
}