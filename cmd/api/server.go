package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/flow"
)

func (app *application) server() *http.Server {
	return &http.Server{
		Addr:         ":80",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (app *application) routes() http.Handler {
	mux := flow.New()

	mux.Use(app.middlewares.recoverPanic)
	mux.Use(app.middlewares.loggedUser)
	mux.Use(app.middlewares.allowCors)

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.HandleFunc("/status", app.status, http.MethodGet)

	mux.HandleFunc("/v1/event/:token", app.handleEvent, http.MethodPost)

	return mux
}
