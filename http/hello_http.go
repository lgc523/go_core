package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	req  = http.Request{}
	resp = http.Request{}
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world!")
	}
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/time/", func(w http.ResponseWriter, req *http.Request) {
		t := time.Now()
		tm := fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.WriteHeader(201)
		w.Write([]byte(tm))
	})
	http.ListenAndServe(":8989", nil)
}
