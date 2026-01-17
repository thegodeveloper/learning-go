package arrays

import "fmt"

// PrintType demonstrates printing array type information
func PrintType(show bool) {
	if show {
		fmt.Println("\n-- Print array type")
		q := [...]int{1, 2, 3, 4, 5, 6, 7}

		fmt.Printf("The array q has the type: %T\n", q)
		fmt.Println("The array values are:", q)
	}
}
