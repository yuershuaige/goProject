package _interface

import "testing"

type Programmer interface {
	WriteHelloWorld () string
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() string {
	 return "fmt.Println(\"helloWorld\")"
}

func TestClient(t *testing.T) {
	var gp Programmer
	gp = new(GoProgrammer)
	t.Log(gp.WriteHelloWorld())
}





