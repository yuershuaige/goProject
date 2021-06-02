package type_test

import "testing"

type Mytype int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64 = 2
	// b = a // 隐式普通类型转换不可以
	b = int64(a) // 显式类型转换
	var c Mytype
	c = Mytype(b)
	t.Log(a,b,c)
}

func TestPoiont(t *testing.T){
	a := 1
	aPtr := &a // 使用取址符获取地址
	//aPtr = aPtr + 1 // goLang不支持指针运算
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)
}

func TestString(t *testing.T){
	// go中 string是值类型而不是引用类型  所以默认值是""而不是null
	var a string
	t.Log("*" + a + "*")
	t.Log(len(a))
	if a == ""{
		t.Log("is blank")
	}
}
