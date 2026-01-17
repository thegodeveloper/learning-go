package forloops

import "fmt"

// RangeIsACopy demonstrates that range creates copies of values
func RangeIsACopy(show bool) {
	if show {
		fmt.Println("--- Range is a copy")

		evenVals := []int{2, 4, 6, 8, 10, 12}

		for _, v := range evenVals {
			v *= 2
		}

		fmt.Println("evenVals:", evenVals)
	}
}
