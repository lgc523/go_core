package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := Service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	//ret := <-AsyncService()
	//t.Logf(ret)
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 30):
		t.Error("time out")
	}
}
