package main

import (
	"fmt"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
	"html"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Created")
}

func main() {
	// get Mongo Session
	session, err := mgo.Dial("mongodb://collector:Ci1aTh1ooshiib6iepha4oongaeSho@mongo/stats")
	// session, err := mgo.Dial("mongodb://collector:Ci1aTh1ooshiib6iepha4oongaeSho@localhost/stats")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(middleware.RedirectSlashes)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", index)

	// RESTy routes for "data" resource
	r.Route("/data", func(r chi.Router) {
		r.Get("/", index)
		r.Post("/", createRecord)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
