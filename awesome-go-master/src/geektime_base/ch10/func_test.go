package ch10

import (
	"fmt"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return 1, 2
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSf := timeSpentFunc(slowFun)
	t.Log(tsSf(10))
}

func timeSpentFunc(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 3, 4, 5, 6, 7, 8))
	t.Log(sum(1, 3, 4, 5, 6, 7, 8, 23423, 234, 2))
}
func clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer clear()
	fmt.Println("start")
	panic("err")
}
