package goroutine_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { mut.Unlock() }()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafeWait(t *testing.T) {
	var mut sync.Mutex
	var w sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			w.Add(1)
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			w.Done()
		}()
	}
	w.Wait()
	t.Logf("counter = %d", counter)
}
