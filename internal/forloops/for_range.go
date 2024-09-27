package forloops

import "fmt"

func forRange(show bool) {
	if show {
		fmt.Println("--- for range")

		evenVals := []int{2, 4, 6, 8, 10, 12}
		for i, v := range evenVals {
			fmt.Println(i, v)
		}
	}
}
