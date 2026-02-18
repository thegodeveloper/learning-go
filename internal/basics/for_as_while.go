package basics

import "fmt"

func ForASWhile(show bool) {
	if show {
		i := 1

		for i < 10 {
			fmt.Println("Iteration:", i)
			i++
		}
	}
}
