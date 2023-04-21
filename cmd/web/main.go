package main

import (
	"log"
	"net/http"
)

func main() {
	routes := routes()
	log.Println("Starting server on :5000")

	// Start the server
	_ = http.ListenAndServe(":5000", routes)
}
