package slices

import "fmt"

func sliceWithMake(show bool) {
	if show {
		x := make([]int, 7)
		fmt.Println("--- Slice with make ---")
		fmt.Println("x value:", x)

		fmt.Println("--- Setting the initial capacity with make ---")
		y := make([]int, 7, 10)
		fmt.Println("y values:", y)
		fmt.Println("y length:", len(y))
		fmt.Println("y capacity:", cap(y))
	}
}
