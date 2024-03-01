package warrays

import "fmt"

func comparingArrays(show bool) {
	if show {
		fmt.Println("\n-- Comparing Arrays")

		a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}

		fmt.Println("a == b:", a == b)
		fmt.Println("a == c:", a == c)
		fmt.Println("b == c:", b == c)
	}
}
