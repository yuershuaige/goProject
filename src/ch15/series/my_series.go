package series

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func Square(n int) int {
	return n * n
}
func GetFibonacciSerie(n int) []int {
	res := []int{1, 1}
	for i := 2; i < n; i++ {
		res = append(res, res[i-2]+res[i-1])
	}
	return res
}
