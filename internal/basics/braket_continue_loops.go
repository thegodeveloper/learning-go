package basics

import "fmt"

func BreakContinueLoops(show bool) {
	if show {
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				continue
			}
			fmt.Println("Odd number", i)
			if i == 5 {
				break
			}
		}
	}
}
