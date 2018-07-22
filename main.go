package main

import (
	"log"

	"flag"
	"labix.org/v2/mgo"
	"net/http"
)

var mongoURL string

//initialize application
func initialize() {
	//	parse flags
	flag.StringVar(&mongoURL, "mongoUrl", "mongodb://localhost/stats", "db url")
	flag.Parse()
}

func main() {
	initialize()
	log.Printf("starting with db url %q", mongoURL)

	// get Mongo Session
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	router := newRouter(session)

	log.Fatal(http.ListenAndServe(":8080", router))
}
