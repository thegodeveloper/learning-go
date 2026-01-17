package functions

import "fmt"

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// VariadicParameters demonstrates variadic function parameters
func VariadicParameters(show bool) {
	if show {
		fmt.Println("--- Variadic Parameters ---")

		fmt.Println(addTo(3))
		fmt.Println(addTo(3, 2))
		fmt.Println(addTo(3, 5, 7, 11, 13, 17))
		a := []int{4, 3}
		fmt.Println(addTo(7, a...))
		fmt.Println(addTo(7, []int{1, 3, 5, 7, 11, 13, 17}...))
	}
}
