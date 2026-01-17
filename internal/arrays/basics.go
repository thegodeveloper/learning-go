package arrays

import "fmt"

// Basics demonstrates basic array operations
func Basics(show bool) {
	if show {
		fmt.Println("\n-- Arrays Basics")
		// this creates an array of 5 items, and each item is initialized to 0
		var a [5]int
		fmt.Println(a)
		fmt.Printf("a length: %d\n", len(a))

		// this creates an array of 5 items, but each item is initialized to a specific value
		b := [5]int{0, 1, 2, 3, 4}
		fmt.Println(b)
		fmt.Printf("b length: %d\n", len(b))

		// set the values for the 3 first items
		c := [5]int{0, 1, 2}
		fmt.Println(c)
		fmt.Printf("c length: %d\n", len(c))

		// sparse array: an array where most elements are set to their zero value
		x := [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("x value: ", x)
	}
}
