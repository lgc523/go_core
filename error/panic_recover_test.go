package error_test

import (
	"errors"
	"fmt"
	"testing"
)

//os.Exit 退出时不会调用 defer，输出调用栈

func TestPanicVsExit(t *testing.T) {
	fmt.Println("start")
	defer func() {
		fmt.Println("finally")
	}()
	defer func() {
		if err := recover(); err != nil {
			//recover
			fmt.Println("recover from :", err)
			//doSomething
		}
	}()
	panic(errors.New("SOMETHING WRONG"))
	//os.Exit(-1)
}
