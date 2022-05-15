package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton *struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("addr:%x\n", unsafe.Pointer(GetSingletonObj()))
			wg.Done()
		}()
	}
	wg.Wait()
}
