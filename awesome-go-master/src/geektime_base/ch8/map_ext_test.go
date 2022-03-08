package ch8

import "testing"

// NOTE: in order to see log, run in vscode with Green Triangle. not `run test` link.
// 极客时间-GO- 11 Map与工厂模式，在Go语言中实现Set

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	// 2 is argument
	t.Log(m[1](2))
	t.Log(m[2](2))
	t.Log(m[3](2))
}

func TestMapForSet(t *testing.T) {
	// impl a set in go
	mySet := map[int]bool{}
	mySet[1] = true

	n := 2
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)

	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet))

}
