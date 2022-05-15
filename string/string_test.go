package string_test

import "testing"

//string 是只读的 byte slice,len 返回包含的byte数
//string byte 可以存储任何数据

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	s = "中"
	t.Log(len(s))
	//string->unicode slice
	c := []rune(s)
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

	//Unicode code point 字符编码集
	//UTF8 是 unicode 存储实现，转换字节序列的规则
}

//strings
//strconv
