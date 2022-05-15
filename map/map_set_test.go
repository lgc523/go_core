package map_test

import (
	"testing"
)

func TestMap(t *testing.T) {
	var m = map[int]int{}
	m[1] = 23
	t.Log(m[1])
	m1 := make(map[string]int, 10)

	if v, ok := m1["23"]; ok {
		t.Logf("key:%s absent", "23")
	} else {
		t.Logf("key:%s is present,val=%d", "23", v)
	}
}

//map value 可以是方法

type Op func(int) int

func TestMapWithFunValue(t *testing.T) {
	m := map[int]Op{}
	m[1] = func(i int) int {
		return i
	}
	m[2] = func(i int) int {
		return i * i
	}
	m[3] = func(i int) int {
		return i * i * i
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	s := map[int]bool{}
	s[1] = true
	n := 2
	delete(s, n)
	if _, ok := s[n]; ok {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(s))
}
