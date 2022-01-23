package dependency_injection

// -- refactoring ioc.go

import "errors"

// TODO: 这里依赖的是其实是一个协议，这个协议是一个没有参数的函数数组

// Undo 本来就是一个类型，不必是一个结构体，是一个函数数组也没有什么问题
type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

// TODO: Undo extract 'Control Logic' from UndoableIntSet

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

/*
TODO: compare with ioc.go
这个就是控制反转，不是由控制逻辑 Undo 来依赖业务逻辑 IntSet，[1]
而是由业务逻辑 IntSet 依赖 Undo [2]

这里依赖的是其实是一个协议，这个协议是一个没有参数的函数数组。
可以看到，这样一来，我们 Undo 的代码就可以复用了
*/

type IntSet struct {
	data map[int]bool
	undo Undo // <- new [2]
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		// new
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) Undo() error {
	// new
	return set.undo.Undo()
}
