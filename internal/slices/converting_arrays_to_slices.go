package slices

import "fmt"

func ConvertingArraysToSlices(show bool) {
	if show {
		fmt.Println("--- Converting array to slice ---")

		x := [4]int{1, 2, 3, 4}
		y := x[:]

		fmt.Println("x values:", x)
		fmt.Println("y values:", y)

		fmt.Printf("x type: %T array\n", x)
		fmt.Printf("y type: %T slice\n", y)

		// Changing a slice element
		y[0] = 7
		fmt.Println("y values:", y)
	}
}
