package slices

import "fmt"

func AppendOverlappingSlices(show bool) {
	if show {
		fmt.Println("--- append makes overlapping slices more confusing ---")

		x := []int{1, 2, 3, 4}
		y := x[:2]
		fmt.Printf("x cap: %v, y cap: %v\n", cap(x), cap(y))
		fmt.Printf("x: %v, y: %v\n", x, y)

		y = append(y, 30)

		fmt.Println("x:", x)
		fmt.Println("y:", y)
	}
}
