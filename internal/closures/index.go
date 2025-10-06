package closures

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("--- Closures ---")

		declaration(false)
	}
}
