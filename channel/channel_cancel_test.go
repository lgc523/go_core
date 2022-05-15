package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func isCancel(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}
func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancel(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
				fmt.Println(i, "Done")
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChan)
	}
	time.Sleep(time.Millisecond * 5)
	cancel2(cancelChan)
	time.Sleep(time.Second)
}
