package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"gopkg.in/mgo.v2"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q. Collection : %q\n", html.EscapeString(r.URL.Path))
}

//createHandler handles records creation
func createHandler(w http.ResponseWriter, r *http.Request, collection *mgo.Collection) {

	record, err := newRecord(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Creating record: %+v", record)

	insertErr := collection.Insert(record)
	if insertErr != nil {
		http.Error(w, insertErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func newRecord(r *http.Request) (*record, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 20*1024*1024)) // Let's read max 20 Mb
	if err != nil {
		return nil, err
	}

	if err := r.Body.Close(); err != nil {
		return nil, err
	}

	// will try to parse data. If it is valid json save parsed
	var parsed interface{}
	parseErr := json.Unmarshal(body, &parsed)

	if parseErr != nil {
		parsed = nil
	}

	record := &record{
		Timestamp: time.Now().UnixNano(),
		Raw:       body,
		String:    string(body),
		Data:      parsed,
	}

	return record, nil
}

//makeHandler helps to pass mongo session to handle and makes sure that this is a copy of session
func makeHandler(fn func(http.ResponseWriter, *http.Request, *mgo.Collection), session *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// make new session
		newSession := session.Copy()
		defer newSession.Close()

		collectionName := chi.URLParam(r, "collectionName")

		if collectionName == "" {
			collectionName = "default"
		}

		collection := newSession.DB("").C(collectionName)

		fn(w, r, collection)
	}
}
