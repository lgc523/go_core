package context_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type otherContext struct {
	context.Context
}

func TestContextWork(t *testing.T) {
	root := context.TODO()
	//a
	ctxA, cancel := context.WithCancel(root)
	go work(ctxA, "work_a")

	//b
	ctxB, _ := context.WithDeadline(ctxA, time.Now().Add(time.Second*3))
	go work(ctxB, "work_b")

	//c
	ctxC := context.WithValue(otherContext{ctxB}, "key", "val")
	go workWithVal(ctxC, "work_c")
	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 5)
	fmt.Println("main over.")
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel.\n", name)
			return
		default:
			fmt.Printf("%s is running.\n", name)
			time.Sleep(time.Second)
		}
	}
}
func workWithVal(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel.\n", name)
			return
		default:
			value := ctx.Value("key")
			decVal := value.(string)
			fmt.Printf("%s is running val=%s\n", name, decVal)
			time.Sleep(time.Second)
		}
	}
}
