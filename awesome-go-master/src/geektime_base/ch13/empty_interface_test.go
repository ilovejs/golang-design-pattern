package ch13

import (
	"fmt"
	"testing"
)

func DoSomeThing(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if j, ok := p.(string); ok {
		fmt.Println("string", j)
		return
	}
	fmt.Println("Unknown Type")
}

func DoSomeThingWay(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomeThing(10)
	DoSomeThing("10")
	DoSomeThingWay(10)
	DoSomeThingWay("10")
}
