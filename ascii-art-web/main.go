package main

import (
	"fmt"
	"log"
	"net/http"

	backend "ascii-art/Backend"
	errorhandling "ascii-art/Error"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", backend.Webhandler)
	mux.HandleFunc("/submit", backend.HandleRequest)
	mux.HandleFunc("/badrequest", errorhandling.BadRequestHandler)
	mux.HandleFunc("/notFound", errorhandling.NotFoundHandler)

	fmt.Println("Server running on http://localhost:5500/")

	err := http.ListenAndServe(":5500", mux)
	log.Fatal(err)
}
