package decoration

import (
	"fmt"
	"reflect"
)

/**
对于 Go 的修饰器模式，还有一个小问题，
那就是好像无法做到泛型。比如上面那个计算时间的函数，其代码耦合了需要被修饰的函数的接口类型，
无法做到非常通用。如果这个问题解决不了，那么，这个修饰器模式还是有点不好用的
*/

// Decorator 第一个是出参 decoPtr ，就是完成修饰后的函数；
// 第二个是入参 fn ，就是需要修饰的函数。
func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

func Run() {
	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	_ = Decorator(&myfoo, foo)
	myfoo(1, 2, 3)

	// OR if u don't want extra types
	mybar := bar
	_ = Decorator(&mybar, bar)
	mybar("hello,", "world!")
}
