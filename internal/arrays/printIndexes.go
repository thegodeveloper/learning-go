package arrays

import "fmt"

func printIndexesAndElements(show bool) {
	if show {
		fmt.Println("\n-- Print Indexes and Elements")
		var a [7]int

		for i, v := range a {
			fmt.Printf("%d %d\n", i, v)
		}
	}
}
