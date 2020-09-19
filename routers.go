package main

import (
	"gopkg.in/mgo.v2"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

//newRouter creates chi router and pass mongo session to handlers
func newRouter(session *mgo.Session) *chi.Mux {
	// Configure http routing
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge: int(24 * time.Hour),
	}))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// RESTy routes for "data" resource
	r.Route("/{collectionName}", func(r chi.Router) {
		r.Get("/", indexHandler)
		r.Post("/", makeHandler(createHandler, session))
	})

	return r
}
