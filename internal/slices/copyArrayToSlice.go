package slices

import "fmt"

func copyArrayToSlice(show bool) {
	if show {
		fmt.Println("--- copy array to slice ---")

		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[2:]
		x[0] = 10

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	}
}
