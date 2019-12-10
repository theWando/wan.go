package routes

import (
	"log"
	"net/http"
)

type UrlMapping struct{}

func (UrlMapping) ServeHTTP(_ http.ResponseWriter, req *http.Request) {
	log.Println("Request received! using ", req.Method)

}
