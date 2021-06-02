package func_variable_defer_test

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	Sum(1,2,3,4,5)
}

func Sum(ops ... int) {
	res := 0
	for _, op := range ops{
		res += op
	}
	fmt.Println(res)
}

func TestDefer (t *testing.T) {
	// 匿名函数  defer类型表示延迟执行  方法的最后执行  类似于java中的finally语句块 用于资源释放
	defer func() {
		fmt.Println("Clear resource.")
	}()
	fmt.Println("Start")
	// 类似于java中的exception 即使出现异常错误 在本方法的最后也会执行defer函数
	panic("err")
	// 不可达的代码段
	fmt.Println("End")
}

