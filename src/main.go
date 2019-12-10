package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting application")
	startServer()
}

func startServer() {
	server := loadServerConfiguration()
	log.Fatal(server.ListenAndServe())
}

type UrlMapping struct {
}

func loadServerConfiguration() *http.Server {
	requestHandler := UrlMapping{}
	return &http.Server{
		Addr:           ":8081",
		Handler:        requestHandler,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (u UrlMapping) ServeHTTP(http.ResponseWriter, *http.Request) {
	log.Println("Request received!")
}
