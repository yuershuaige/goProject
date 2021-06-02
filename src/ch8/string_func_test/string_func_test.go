package string_func_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestFuncTest(t *testing.T) {
	s := "a,b,f,g,h,j,k"
	parts := strings.Split(s, ",")
	for _,part := range parts {
		t.Logf(part)
	}
	t.Log(strings.Join(parts,"-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	// t.Log("str" + 10) 直接写字符串+整形值 会编译错误
	t.Log("str" + s)
	if i,err := strconv.Atoi("20"); err == nil {
		t.Log(5 + i)
	}
}
