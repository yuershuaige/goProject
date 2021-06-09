package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError error = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibSlice := []int{1, 1}
	for i := 2; i < n; i++ {
		fibSlice = append(fibSlice, fibSlice[i-2], fibSlice[i-1])
	}
	return fibSlice, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("it is less ..")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

/**
避免类似这种的错误嵌套
*/
func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

/**
遇到错误应该快速失败
*/
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(list)
}
