package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p * Pet) speak () {
	fmt.Print("...")
}

func (p * Pet) speakTo(host string) {
	p.speak()
	fmt.Print(" ", host)
}

type Dog struct {
	//p *Pet
	// 匿名嵌套类型 而这种内嵌的类型不支持方法重载 不会调用子类复写的方法 所以不支持lsp
	Pet
}

func (d * Dog) speak () {
	fmt.Print("wang wang wang")
}
//
//func (d * Dog) speakTo(host string) {
//	d.speak()
//	fmt.Print(" ", host)
//}


func TestDog(t *testing.T) {
	//此处和java不同 golang不支持lsp
	//cannot use new(Dog) (type *Dog) as type Pet in assignment
	//var dog Pet = new(Dog)
	dog := new(Dog)
	dog.speakTo("chao")
}
