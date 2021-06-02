package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T){
	//var prev int = 1
	//var curr int = 1
	//var (
	//	prev int = 1
	//	// 类型推断
	//	curr  = 1
	//)
	// 类型推断
	prev := 1
	curr := 1
	fmt.Print(prev)
	for i := 0; i < 5; i ++ {
		fmt.Print(" ",curr)
		temp := curr
		curr = prev + curr
		prev = temp
	}
	fmt.Println(" ")
}

func TestExchange(t *testing.T){
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	// 等价
	a,b = b,a
	t.Log(a,b)
}