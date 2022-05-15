package interface_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent:%f seconds", time.Since(start).Seconds())
		return ret
	}
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func slowFun(op int) int {
	time.Sleep(time.Second)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSf := timeSpent(slowFun)
	t.Log(tsSf(10))
}
