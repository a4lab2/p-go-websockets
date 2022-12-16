package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	m := pat.New()

	m.Get("/homepage", http.HandlerFunc(app.homepage))

	m.Get("/ws", http.HandlerFunc(app.websocket))

	return m
}
