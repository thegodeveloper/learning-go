package pointers

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("--- Pointers ---")

		declaration(false)
	}
}
