package main

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	//restful
	//resource oriented architecture

	//prefix tree
	router := httprouter.New()
	// router.GET("/", IndexRouter)
	router.GET("/hello/:name", HelloRouter)
	router.GET("/fb", FbRouter)
	log.Fatal(http.ListenAndServe(":8989", router))
}

func IndexRouter(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func HelloRouter(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello,%s!\n", ps.ByName("name"))
}

func FbRouter(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fibonacci, _ := getFibonacci(23)
	w.Write([]byte(fmt.Sprintf("%v", fibonacci)))
}

var illegalErr = errors.New("N should be in [2,100")

func getFibonacci(n int) ([]int, error) {
	if n < 0 || n > 100 {
		return nil, illegalErr
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
