package main

import (
	"github.com/theWando/wan.do/src/conf"
	"github.com/theWando/wan.do/src/routes"
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

func loadServerConfiguration() *http.Server {
	requestHandler := routes.UrlMapping{}
	return &http.Server{
		Addr:           ":" + conf.GetPort(),
		Handler:        requestHandler,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
