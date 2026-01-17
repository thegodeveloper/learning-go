package arrays

import "fmt"

// PrintIndexesAndElements demonstrates printing array indexes and values
func PrintIndexesAndElements(show bool) {
	if show {
		fmt.Println("\n-- Print Indexes and Elements")
		var a [7]int

		for i, v := range a {
			fmt.Printf("%d %d\n", i, v)
		}
	}
}
