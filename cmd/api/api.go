package main

import (
	"log"
	"net/http"
	"time"

	"github.com/faustcelaj/social_project/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	// config the db into the runtime so each environment has access to it
	db  dbConfig
	env string
}

type dbConfig struct {
	// connection address string for the PostgreSQL database
	addr string
	// Specifies the maximum number of open connections allowed to the database
	// Limiting this helps control database load and resource usage
	maxOpenConns int
	// Defines the maximum number of idle connections
	// helps reduce the number of active connections that are not being used, which saves resources
	maxIdleConns int
	// Specifies the maximum idle time allowed for a connection
	maxIdleTime string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// we can group multiple routes together and set what we want to do with whatever comes after
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		// if server takes more than 30 sec to write we timeout
		WriteTimeout: time.Second * 30,
		// if the client takes more than 10 seconds to read our response we timeout
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
