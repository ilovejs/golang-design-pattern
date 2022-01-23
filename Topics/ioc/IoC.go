package ioc

import "errors"

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}
func (set *IntSet) Add(x int) {
	set.data[x] = true
}
func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

/** TODO: 我们在 UndoableIntSet 中嵌入了IntSet ，然后 Override 了 它的 Add()和 Delete()
  Contains() 方法没有 Override，所以，就被带到 UndoableInSet 中来了
*/

type UndoableIntSet struct { // Poor style
	IntSet    // Embedding (delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

// Add 在 Override 的 Add()中，记录 Delete 操作
func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		// Delete coming from IntSet
		set.functions = append(set.functions, func() { set.Delete(x) }) // [1] see dependency_inversion.go
	} else {
		set.functions = append(set.functions, nil)
	}
}

// Delete 在 Override 的 Delete() 中，记录 Add 操作
func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
		delete(set.data, x)
		// Add coming from IntSet
		set.functions = append(set.functions, func() { set.Add(x) }) // [1] see dependency_inversion.go
	} else {
		set.functions = append(set.functions, nil)
	}
}

/*
[pros]用这样的方式为已有的代码扩展新的功能是一个很好的选择。
这样，就可以在重用原有代码功能和新的功能中达到一个平衡。

[cons]但是，这种方式最大的问题是，
TODO: Undo 操作其实是一种控制逻辑，并不是业务逻辑，
所以，在复用 Undo 这个功能时，是有问题的，

因为其中加入了大量跟 IntSet 相关的业务逻辑
*/

// Undo is new method
func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}
