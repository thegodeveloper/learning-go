package forloops

import "fmt"

// ForRange demonstrates for-range loops on slices
func ForRange(show bool) {
	if show {
		fmt.Println("--- for range")

		evenVals := []int{2, 4, 6, 8, 10, 12}
		for i, v := range evenVals {
			fmt.Println(i, v)
		}
	}
}
