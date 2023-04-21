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

	fileServe := http.FileServer(http.Dir("./html/static/"))
	mux.Get("/html/static/", http.StripPrefix("/html/static/", fileServe))

	return mux
}
