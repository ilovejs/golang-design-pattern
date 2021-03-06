package decoration

import "fmt"

// high order function
func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func RunDemo() {
	// use
	decorator(Hello)("Hello, World!")
}
