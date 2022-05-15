package context_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second)
}
