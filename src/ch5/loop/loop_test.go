package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n ++
	}
}

func TestForLoop(t *testing.T) {
	for n := 0; n < 5; n ++ {
		t.Log(n)
	}

	// go和java不同  forloop中加括号会编译错误
	//for (i := 10; i < 15; i ++) {
	//
	//}
}


func TestWhileTrue(t *testing.T) {
	for {
		t.Log("always loop")
	}
}
