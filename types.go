package main

type record struct {
	Timestamp int64       `json:"timestamp"`
	Raw       []byte      `json:"raw"`
	String    string      `json:"string"`
	Data      interface{} `json:"data"`
}
