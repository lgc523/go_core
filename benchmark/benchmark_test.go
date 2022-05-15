package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, e := range elems {
		ret += e
	}
	assert.Equal("12345", ret)
}

func TestConcatStringBybytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, e := range elems {
		buf.WriteString(e)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, e := range elems {
			ret += e
		}
	}
	b.StopTimer()
}
func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, e := range elems {
			buf.WriteString(e)
		}
	}
	b.StopTimer()
}

// go test -bench=. -benchmem
