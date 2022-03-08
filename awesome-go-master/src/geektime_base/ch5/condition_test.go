package ch5

import "testing"

func TestForSyntax(t *testing.T) {
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}

	// 无限循环
	n = 0
	for {
		if n > 5 {
			break
		}
		t.Log(n)
		n++
	}
}

func TestIfSyntax(t *testing.T) {
	if n := somefunc(); n {
		t.Log(n)
	}
}

func somefunc() bool {
	return 1 == 1
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		}
	}
}
