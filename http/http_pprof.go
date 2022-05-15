package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	//go tool pprof http://localhost:8923/debug/pprof/profile
	mux.HandleFunc("/", Index)
	//curl  -X POST -H"Content-type:application/x-www-form-urlencoded"  http://localhost:8923/hello -d "name"="spider"
	mux.HandleFunc("/hello/", Hello)
	mux.HandleFunc("/fb", Fb)
	server := &http.Server{
		Addr:              ":8923",
		Handler:           mux,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 50,
		ReadHeaderTimeout: time.Second * 5,
	}
	log.Println("server is running, listen:", server.Addr[1:])
	log.Fatal(server.ListenAndServe())
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Path:%s,Method:%s", req.URL.Path, req.Method)
		next.ServeHTTP(w, req)
	})
}

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprintf(w, "hello,%s!\n", name)
}

func Fb(w http.ResponseWriter, req *http.Request) {
	fibonacci, _ := GetFibonacci(50)
	w.Write([]byte(fmt.Sprintf("%v", fibonacci)))
}

var illErr = errors.New("N should be in [2,100")

func GetFibonacci(n int) ([]int, error) {
	if n < 0 || n > 100 {
		return nil, illErr
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
