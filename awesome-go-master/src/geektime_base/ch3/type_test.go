package ch3

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a, b)

	// 即使别名也不支持隐式转换，只能直接转换
	var c MyInt
	c = MyInt(b)
	t.Log(b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// Go 语言不支持指针运算

	t.Log(a, aPtr)

	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// Go 字符串默认是空字符，而不是空
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
