package main

import (
	"github.com/theWando/wan.do/src/routes"
	"log"
	"net/http"
	"os"
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
	port := getPort()
	return &http.Server{
		Addr:           ":" + port,
		Handler:        requestHandler,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Using port ", port)
	return port
}
