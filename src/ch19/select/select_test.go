package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		res := service()
		fmt.Println("returned result")
		retCh <- res
		fmt.Println("service exited")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case res := <-AsyncService():
		t.Log(res)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
