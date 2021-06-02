package operator_test

import "testing"

func TestCompareArray(t *testing.T){
	a := [...]int{1,2,3,4,5}
	b := [...]int{1,3,4,5,2}
	c := [...]int{1,3,5,6}
	d := [...]int{1,3,5,6}
	t.Log(a == b)
	//t.Log(a == c) // 长度相等才可以比较
	t.Log(c == d)
}

// 按位 置为0
const(
	Readable = 1 << iota
	Writable
	Executable
)

func  TestConstantTry(t *testing.T){
	a := 7 // 0111
	a = a &^ Readable // 把a上与Readable上位为1的位为0
	t.Log(a & Readable == Readable," ",a & Writable == Writable," ", a & Executable == Executable)
}
