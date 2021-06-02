package condition

import "testing"

func TestIf (t *testing.T){
	if 1 == 1 {
		t.Log("1 == 1")
	}

	// go 和 java不同 if条件不推荐加括号  但是加了也不会导致编译错误
	if ( 1 == 1 ){
		t.Log(" (1 == 1) ")
	}
}

func TestMultiSec(t *testing.T){
	if v,err := someFun();err == nil{
		t.Log("1",v)
	}else {
		t.Log("0",err)
	}
}

func someFun() (interface{}, interface{}) {
	return 1,2
}