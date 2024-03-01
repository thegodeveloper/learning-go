package slices

import "fmt"

func slicingSlices(show bool) {
	if show {
		fmt.Println("--- slicing slices ---")

		// this is a slice
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:2]
		d := x[1:3]
		e := x[:]

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Println("d:", d)
		fmt.Println("e:", e)
	}
}
