package generic

import (
	"fmt"
	"testing"
)

//declare a type constraint

type Number interface {
	int64 | float64
}

func TestSimpleFunc(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		sumInts(ints),
		sumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		sumIntsOrFloats[string, int64](ints),
		sumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums omit parameterType,compiler inferred: %v and %v\n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with interface constraint: %v and %v\n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var f float64
	for _, v := range m {
		f += v
	}
	return f
}

func sumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var result V
	for _, v := range m {
		result += v
	}
	return result
}
