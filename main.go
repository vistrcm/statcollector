package main

import (
	"fmt"
	"log"

	"flag"
	"gopkg.in/mgo.v2"
	"html"
	"net/http"
)

var mongoUrl string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Created")
}

//initialize application
func initialize() {
	//	parse flags
	flag.StringVar(&mongoUrl, "mongoUrl", "mongodb://localhost/stats", "db url")
	flag.Parse()
}

func main() {
	initialize()
	log.Printf("starting with db url %q", mongoUrl)

	// get Mongo Session
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	router := newRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
