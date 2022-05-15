package concurrency

import (
	"runtime"
	"testing"
	"time"
)

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := RunTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

// go test -v first_response_test.go all_response_test.go -run="TestAllResponse"
// not graceful
func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second)
	t.Log("Before:", runtime.NumGoroutine())
}
