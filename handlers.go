package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//createHandler handles records creation
func createHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 20*1024*1024)) // Let's read max 20 Mb
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	log.Printf("Create record: %+v", record)

	//TODO: save record here

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
