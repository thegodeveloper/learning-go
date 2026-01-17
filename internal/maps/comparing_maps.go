package maps

import (
	"fmt"
	"maps"
)

func ComparingMaps(show bool) {
	if show {
		fmt.Println("--- Comparing Maps")

		m := map[string]int{
			"hello": 7,
			"world": 77,
		}
		n := map[string]int{
			"world": 77,
			"hello": 7,
		}

		fmt.Println(maps.Equal(m, n))
	}
}
