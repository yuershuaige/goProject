package polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Print(\"Hello world!\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello world!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T,%v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	g := &GoProgrammer{}
	j := new(JavaProgrammer)
	writeFirstProgram(g)
	writeFirstProgram(j)
}
