package iterator

import "fmt"

func ExampleIterator() {
	var aggregate Aggregate
	aggregate = NewNumbers(1, 10)

	IterPrint(aggregate.Iterator())

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10

	fmt.Print("0")
}
