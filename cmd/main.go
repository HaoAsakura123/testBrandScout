package main

import (
	"log"
	"net/http"
	"testbrandscout/internal/handler"
)

func main() {

	http.HandleFunc("/quotes", handler.HandleQuotes)
	http.HandleFunc("/quotes/", handler.HandleDelete)
	http.HandleFunc("/quotes/random/", handler.HandleRandom)

	log.Println("server is running on port:8080")
	http.ListenAndServe(":8080", nil)
}
