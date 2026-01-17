package slices

import "fmt"

func SlicesWithOverlappingStorage(show bool) {
	if show {
		fmt.Println("Slices with Overlapping Storage")

		// changing x modified both y and z, while changes to y and z modified x.
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
