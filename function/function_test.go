package function

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// timeSpent
// 计算函数执行时间
func timeSpent(inner func(int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent:%f seconds", time.Since(start).Seconds())
		return ret
	}
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
