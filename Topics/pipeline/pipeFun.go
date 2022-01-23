package pipeline

import "fmt"

// 同样，如果你不想有那么多的函数嵌套，就可以使用一个代理函数来完成

// read more about function types https://go101.org/article/function.html
// https://medium.com/@kandros/go-function-type-reusable-function-signatures-2389f6bdd4f6

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

// no need actually
var (
	odd2 = odd // of PipeFunc type
)

func RunDelegate() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//for n := range pipeline(nums, gen, odd, sq, sum) {
	for n := range pipeline(nums, echo, odd2, sq, sum) {
		fmt.Println(n)
	}
}
