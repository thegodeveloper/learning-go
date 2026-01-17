package slices

import "fmt"

func SlicingSlices(show bool) {
	if show {
		fmt.Println("--- slicing slices ---")

		// this is a slice
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:2]
		d := x[1:3]
		e := x[:]

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Println("d:", d)
		fmt.Println("e:", e)
	}
}

func ChangeSlicing(show bool) {
	if show {
		fmt.Println("--- Change Slicing")

		x := []string{"a", "b", "c", "d"}
		y := x[:2]
		z := x[1:]
		x[1] = "y"
		y[0] = "x"
		z[1] = "z"
		fmt.Println("x values:", x)
		fmt.Println("y values:", y)
		fmt.Println("z values:", z)
	}
}
