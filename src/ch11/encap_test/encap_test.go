package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)
// 实体
type Employee struct {
	Id string
	Name string
	Age int
}
// 行为
// 在此实例方法被调用时，实例的成员会进行值复制
func (e Employee) String1 () string{
	fmt.Printf("\n Name's address is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 在此实例方法被调用时，实例的成员不会进行值复制 可以避免内存拷贝
func (e *Employee) String2 () string{
	fmt.Printf("\n Name's address is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"01","zhangone",15}
	e1 := &Employee{"01","zhangone",15}
	fmt.Printf("\n Name's address is %x",unsafe.Pointer(&e.Name))
	t.Log(e.String1())
	t.Log(e1.String1())
	t.Log(e.String2())
	t.Log(e1.String2())
}

func TestCreateEmployObj(t *testing.T) {
	e1 := Employee{"01","zhangjie",15}
	e2 := Employee{Name :  "wangwu", Age : 18}
	e3 := new(Employee) // 这里返回的是指针(引用传递) 相当于 e3 := &Employee{}
	// 与其他主流语言不同： 通过实例的指针访问成员不需要使用 ->
	e3.Age = 15
	e3.Id = "02"
	e3.Name = "zhaoliu"
	t.Log(e1)
	// 输出e1的指针
	t.Log(&e1)
	t.Log(e1.Name)
	t.Log(e2)
	t.Log(e2.Name)
	t.Log(e3)
	t.Log(e3.Name)
	t.Logf("e1 is %T", e1)
	t.Logf("e3 is %T", e3)
}
