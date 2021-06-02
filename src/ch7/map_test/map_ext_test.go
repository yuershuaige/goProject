package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	// 函数是 golang 的一等公民
	m := map[int]func(op int) int {}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op + 5
	}
	m[3] = func(op int) int {
		return op * op
	}
	t.Log(m[1](2),m[2](2),m[3](2))
}

func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true
	n := 1
	if set[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not existing",n)
	}
	set[3] = true
	t.Log(len(set))
	delete(set,1)
	t.Log(len(set))
	if set[n] {
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not existing",n)
	}
}

//func TestMapForSetUseInt(t *testing.T) {
//	// int 不可以通过golang的编译 用于布尔判断
//	set := map[int]int{}
//	set[1] = 1
//	n := 1
//	if set[n] {
//		t.Logf("%d is existing",n)
//	} else {
//		t.Logf("%d is not existing",n)
//	}
//	set[3] = 1
//	t.Log(len(set))
//	delete(set,1)
//	t.Log(len(set))
//	if set[n] {
//		t.Logf("%d is existing",n)
//	} else {
//		t.Logf("%d is not existing",n)
//	}
//}