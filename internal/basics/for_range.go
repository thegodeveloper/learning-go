package basics

import "fmt"

func ForRange(show bool) {
	if show {
		for i := range 10 {
			fmt.Println("For loop:", i)
		}
	}
}
