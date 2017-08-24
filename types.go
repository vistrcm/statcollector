package main

import "fmt"

type record struct {
	Timestamp int64
	RawData []byte
	StringData string
	Data interface{}
}


func (r *record) show() {
	fmt.Printf("record. Time: %d. Data %q", r.Timestamp, r.Data)
}
