package test

import (
	"fmt"
	"testing"
)

func square(i int) int {
	return i*i + 1
}
func TestSquare(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	fmt.Printf("%T\n", input)
	fmt.Printf("%T\n", expected)
	for i := 0; i < len(input); i++ {
		ret := square(input[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected is %d,the actual is %d\n", input[i], expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

// fatal interrupt
func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}

//go test -v -cover
//assert https://github.com/stretchr/testify
