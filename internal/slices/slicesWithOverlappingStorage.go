package slices

import "fmt"

func slicesWithOverlappingStorage(show bool) {
	if show {
		fmt.Println("Slices with Overlapping Storage")

		x := []string{"A", "B", "C", "D"}
		y := x[:2]
		z := x[1:]
		x[1] = "Y"
		y[0] = "X"
		z[1] = "Z"
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	}
}
