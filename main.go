package main

import (
	"log"

	"flag"
	"gopkg.in/mgo.v2"
	"net/http"
)

var mongoUrl string

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

	router := newRouter(session)

	log.Fatal(http.ListenAndServe(":8080", router))
}
