package slices

import "fmt"

func MultiDimensional(show bool) {
	if show {
		twoD := make([][]int, 3)
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println(twoD)
	}
}
