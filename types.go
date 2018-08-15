package main

import (
	"mime/multipart"
	"net/http"
	"net/url"
)

type request struct {
	Method           string          `json:"Method"`
	URL              *url.URL        `json:"URL"`
	Proto            string          `json:"Proto"`
	Header           http.Header     `json:"Header"`
	ContentLength    int64           `json:"ContentLength"`
	TransferEncoding []string        `json:"TransferEncoding"`
	Host             string          `json:"Host"`
	Form             url.Values      `json:"Form"`
	PostForm         url.Values      `json:"PostForm"`
	MultipartForm    *multipart.Form `json:"MultipartForm"`
	Trailer          http.Header     `json:"Trailer"`
	RemoteAddr       string          `json:"RemoteAddr"`
	RequestURI       string          `json:"RequestURI"`
	RequestID        string          `json:"RequestID"`
}

type metadata struct {
	Request request `json:"request"`
}

type record struct {
	Timestamp int64       `json:"timestamp"`
	Raw       []byte      `json:"raw"`
	String    string      `json:"string"`
	Data      interface{} `json:"data"`
	Metadata  metadata    `json:"metadata"`
}
