package error_test

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("n must be in the range [0,23]")
	println(err.Error())
}

var IllegalErr = errors.New("N should be in [2,100")

func GetFibonacci(n int) ([]int, error) {
	if n < 0 || n > 100 {
		return nil, IllegalErr
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if fibonacci, err := GetFibonacci(-10); err != nil {
		if err == IllegalErr {
			t.Error(err)
		} else {
			t.Log(fibonacci)
		}
	}
}
