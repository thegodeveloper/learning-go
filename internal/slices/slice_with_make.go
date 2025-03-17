package slices

import "fmt"

func sliceWithMake(show bool) {
	if show {
		x := make([]int, 7)
		fmt.Println("--- Slice with make ---")
		fmt.Println("x value:", x)

		fmt.Println("--- Setting the initial capacity with make ---")
		y := make([]int, 7, 14) // append double the current capacity if there is no room for new values
		fmt.Println("y values:", y)
		fmt.Println("y length:", len(y))
		fmt.Println("y capacity:", cap(y))

		fmt.Println("--- Create a slice with zero length but a capacity that's greater than zero ---")
		z := make([]int, 0, 10)
		fmt.Println("z values:", z)
		fmt.Println("z length:", len(z))
		fmt.Println("z capacity:", cap(z))
		// you cannot directly index into z slice, but you can append values to it
		z = append(z, 10, 20, 30, 40)
		fmt.Println("z values:", z)
		fmt.Println("z length:", len(z))
		fmt.Println("z capacity:", cap(z))
	}
}
