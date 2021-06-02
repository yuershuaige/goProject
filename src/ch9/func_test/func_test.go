package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a,b := returnMultiValues()
	t.Log(a,b)
	c,_ := returnMultiValues()
	t.Log(c)

	tsSF := timeSpend(slowFunc)
	t.Log(tsSF(55))
}

func timeSpend(inner func(op int) int) func(op int) int{
	return func(n int) int {
		startTime := time.Now()
		res := inner(n)
		fmt.Println("time spend:", time.Since(startTime))
		return res
	}
}

func slowFunc(op int) int{
	time.Sleep(time.Second * 1)
	return op
}