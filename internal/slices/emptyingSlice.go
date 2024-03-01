package slices

import "fmt"

func emptyingSlice(show bool) {
	if show {
		fmt.Println("--- Emptying a slice")

		s := []string{"first", "second", "third"}
		fmt.Println("s values:", s)
		fmt.Println("s len:", len(s))
		clear(s)
		fmt.Println("s values after cleaning it:", s)
		// After cleaning a slice, the len doesn't change
		fmt.Println("s len:", len(s))
	}
}
