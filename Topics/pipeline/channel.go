package pipeline

import "fmt"

// 它会把一个整型数组放到一个 Channel 中，并返回这个 Channel
func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

// took from https://go.dev/blog/pipelines
// similar to 'echo'
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

// Client 用户端的代码如下所示（注：你可能会觉得，sum()，odd() 和 sq() 太过于相似，
// 其实，你可以通过 Map/Reduce 编程模式
// 或者是 Go Generation 的方式合并一下
// 上面的代码类似于我们执行了 Unix/Linux 命令： echo $nums | sq | sum
func Client() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for n := range sum(sq(odd(echo(nums)))) {
		fmt.Println(n)
	}
}
