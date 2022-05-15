package interface_test

import (
	"fmt"
	"testing"
)

type Code string
type Programming interface {
	WriteHelloWorld() Code
}
type GoProgramming struct {
}

func (gp *GoProgramming) WriteHelloWorld() Code {
	return "fmt.Println(\"hello world\")"
}

type JavaProgramming struct {
}

func (jp *JavaProgramming) WriteHelloWorld() Code {
	return "System.out.println(\"hello world\")"
}

//interface 参数只能是指针
func WriteFirstProgramming(p Programming) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}
func TestPolymorphism(t *testing.T) {
	g := &GoProgramming{}
	j := new(JavaProgramming)
	WriteFirstProgramming(g)
	WriteFirstProgramming(j)
}
