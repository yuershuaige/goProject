package empty_interface

import (
	"fmt"
	"testing"
)

/**
空接口 p interface{}
断言  p.(type)
*/
func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("String", s)
	//	return
	//}
	switch i := p.(type) {
	case int:
		fmt.Println("Integer", i)
	case string:
		fmt.Println("String", i)
	default:
		fmt.Println("unknow type", i)
	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
