package maps

import "fmt"

func commaOkIdiomInMaps(show bool) {
	if show {
		fmt.Println("--- Comma OK Idiom in Maps")
		m := map[string]int{
			"hello": 7,
			"world": 7,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok)

		v, ok = m["world"]
		fmt.Println(v, ok)

		v, ok = m["goodbye"]
		fmt.Println(v, ok) // if the key doesn't exist in the map, it returns false and assign the zero value to v
	}
}
