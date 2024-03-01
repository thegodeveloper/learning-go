package slices

import "fmt"

func sliceAppend(show bool) {
	if show {
		fmt.Println("--- Slice Append")

		var x []int
		fmt.Println("Initial Slice Length:", len(x))
		fmt.Println("Initial Slice Capacity:", cap(x))
		x = append(x, 10)
		fmt.Println("Slice Append, value of x:", x)
		fmt.Println("Slice Length:", len(x))
		fmt.Println("Slice Capacity:", cap(x))
		x = append(x, 20)
		fmt.Println("Slice Capacity:", cap(x))
		x = []int{10, 20, 30}
		fmt.Println("Current x values:", x)
		fmt.Println("New slice capacity:", cap(x))
	}
}
