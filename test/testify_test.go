package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(12345, 12345, "the should be equal")
}
