package slices

import "fmt"

func confusingSlices(show bool) {
	if show {
		fmt.Println("--- Confusing Slices ---")

		x := make([]string, 0, 5)
		x = append(x, "a", "b", "c", "d")
		y := x[:2]
		z := x[2:]
		fmt.Printf("cap x: %d, cap y: %d, cap z: %d\n", cap(x), cap(y), cap(z))
		y = append(y, "i", "j", "k")
		x = append(x, "x")
		z = append(z, "y")
		fmt.Println("x values:", x)
		fmt.Println("y values:", y)
		fmt.Println("z values:", z)
		fmt.Printf("cap x: %d, cap y: %d, cap z: %d\n", cap(x), cap(y), cap(z))

		// To avoid complicated slice situations, you should either never use append with a sub-slice
		// or make sure that append doesn't cause an overwrite by using a full slice expression.
	}
}
