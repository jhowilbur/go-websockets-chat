package main

import (
	"github.com/bmizerany/pat"
	"github.com/jhowilbur/go-websockets-chat/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	mux.Get("/ws/ping", http.HandlerFunc(handlers.WsPingEndpoint))

	return mux
}
