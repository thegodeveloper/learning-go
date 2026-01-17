package slices

import "fmt"

func CopySlice(show bool) {
	if show {
		fmt.Println("--- copying slice ---")

		x := []int{1, 2, 3, 4}
		y := make([]int, 4)
		num := copy(y, x)

		x = append(x, 7) // slices don't share memory
		x[0] = 77

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("number of elements copied:", num)

		fmt.Println("--- copying a slice is limitted by the smallest one ---")
		z := make([]int, 2)
		num = copy(z, x)
		fmt.Println("z:", z)
		fmt.Println("number of elements copied:", num)
	}
}
