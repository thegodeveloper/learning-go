package forloops

import "fmt"

// ForRangeMaps demonstrates for-range loops on maps
func ForRangeMaps(show bool) {
	if show {
		fmt.Println("--- for range maps")

		m := map[string]int{
			"a": 1,
			"c": 3,
			"b": 2,
		}

		for i := 0; i < 3; i++ {
			fmt.Println("Loop", i)
			for k, v := range m {
				fmt.Println(k, v)
			}
		}
	}
}
