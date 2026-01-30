package basics

import "fmt"

func OuterLoops(show bool) {
	if show {
		rows := 7

		// outer loop
		for i := 1; i <= rows; i++ {
			// inner loop for spaces before stars
			for j := 1; j <= rows-i; j++ {
				fmt.Print(" ")
			}

			// inner loop for stars
			for k := 1; k <= 2*i-1; k++ {
				fmt.Print("*")
			}

			fmt.Println("")
		}
	}
}
