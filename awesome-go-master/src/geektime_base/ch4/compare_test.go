package ch4

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	//a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4, 9}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4, 5}
	e := [...]int{1, 2, 3, 5, 4}

	// 长度不同不可以直接比较
	//t.Log(a == b)
	t.Log(b == c)
	t.Log(c == d)
	// 数组顺序不同，导致结果也不相同
	t.Log(d == e)
}

func TestBitClear(t *testing.T) {
	a := 7
	// 按位清零运算符，主要核心思想，如果右边是零，原先保持不变，如果是1则全部置为零
	a = a &^ Readable
	t.Log(a)
	t.Log(a&Readable == Readable)
}
