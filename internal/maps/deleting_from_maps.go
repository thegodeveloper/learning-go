package maps

import "fmt"

func DeletingFromMaps(show bool) {
	if show {
		fmt.Println("--- Deleting From Maps ---")
		m := map[string]int{
			"hello": 7,
			"world": 7,
		}
		fmt.Println("Original m value:", m)
		delete(m, "world")
		fmt.Println("New m value after removing a key", m)
	}
}
