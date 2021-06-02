package customer_type

import (
	"fmt"
	"testing"
	"time"
)

func TestCustomer(t *testing.T) {
	tsSF := timeSpend(slowFunc)
	t.Log(tsSF(10))
}

type intConv func(op int) int

func timeSpend(inner intConv) intConv{
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


