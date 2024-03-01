package slices

import "fmt"

func copySliceToArray(show bool) {
	if show {
		fmt.Println("--- copying slice to array ---")

		x := []int{1, 2, 3, 4}  // slice
		d := [4]int{5, 6, 7, 8} // array
		y := make([]int, 2)
		copy(y, d[:]) // copying array to slice
		fmt.Println("y:", y)
		copy(d[:], x) // copying slice to array
		fmt.Println("d:", d)
	}
}
