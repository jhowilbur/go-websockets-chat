package main

import (
	"github.com/jhowilbur/go-websockets-chat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()
	go handlers.ListenToWsChannel()

	// Start the server
	log.Println("Starting server on :5000")
	_ = http.ListenAndServe(":5000", mux)
}
