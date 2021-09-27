package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/mike-wb/udemy-go/cloudykit/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/send", http.HandlerFunc(handlers.SendData))

	return mux
}
