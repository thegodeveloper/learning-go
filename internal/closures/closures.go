package closures

import "fmt"

func Index(show bool) {
	if show {
		fmt.Println("--- Closures ---")

		declaration(false)
	}
}
