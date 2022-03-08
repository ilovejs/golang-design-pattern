package ch2_test

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// 方式一
	//var a int = 1
	//var b int = 1

	// 方式二
	//var (
	//	a int = 1
	//	b     = 1
	//)

	// 方式三
	var a = 1
	b := 1

	t.Log(a)
	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

func TestExchange(t *testing.T) {
	a := 1
	b := 3
	tmp := a
	a = b
	b = tmp
	t.Log(a, b)

	a, b = b, a
	t.Log(a, b)

}
