package slices

import "fmt"

func sliceWithSpecificValuesPerIndex(show bool) {
	if show {
		fmt.Println("--- Slice with specific values per index.")

		// this creates a slice of 12 ints
		x := []int{1, 5: 6, 10: 100, 15}
		println("--- slice with specific values per index ---")
		println("x values: ", x)
	}
}
