package slices

import "fmt"

func SliceAppend(show bool) {
	if show {
		fmt.Println("--- Slice Append")

		// This is a nil slice
		var x []int
		fmt.Println("Initial Slice Length:", len(x))
		fmt.Println("Initial Slice Capacity:", cap(x))

		x = append(x, 10)
		fmt.Println("Slice Append, value of x:", x)
		fmt.Println("Slice Length:", len(x))
		fmt.Println("Slice Capacity:", cap(x))

		x = append(x, 20)
		fmt.Println("Slice Capacity:", cap(x))

		x = []int{10, 20, 30}
		fmt.Println("Current x values:", x)
		fmt.Println("New slice capacity:", cap(x))

		// Append more than one value
		x = append(x, 40, 50, 60, 70)
		fmt.Println("New x values:", x)
		fmt.Println("New x capacity:", cap(x))

		// Append a slice into another
		y := []int{80, 90, 100}
		x = append(x, y...)
		fmt.Println("New x values:", x)
		// Because there's room in the slice, append doesn't increase the capacity
		fmt.Println("New x capacity:", cap(x))
	}
}
