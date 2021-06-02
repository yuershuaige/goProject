package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 默认值初始化为""
	t.Log(len(s))
	s = "hello"
	t.Log(s)
	t.Log(len(s)) // 长度是byte长度而不是字符长度 此处刚好一个英文字符占用一个byte
	// s[1] = '3' // string是不可变的slice
	// t.Log(s)
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	t.Log(s)
	t.Log(len(s))
	s = "\xE4\xBA\xBB\xFF" // 即使是无法识别的字符
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(s) // byte数
	t.Log(len(s))
	// 使用rune取出字符串里的Unicode
	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size: ", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x",c[0])
	// utf8是unicode编码的物理存储形式 根据utf8规则将unicode码存储  e4b8ad也就是三个byte
	t.Logf("中 utf8 %x",s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	// range将字符串先通过rune获取unicode编码 再遍历
	for _, c := range s {
		t.Logf("%[1]c %[1]x %[1]d",c)
	}
}
