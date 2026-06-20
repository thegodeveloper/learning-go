package gotousecase

import (
	"fmt"
	"math/rand"
)

func Run(show bool) {
	if show {
		fmt.Println("--- goto use case ---")

		a := rand.Intn(10)
		for a < 100 {
			if a%5 == 0 {
				goto done
			}
			a = a*2 + 1
		}
		fmt.Println("do something when the loop completes normally")
	done:
		fmt.Println("do complicated stuff no matter why we left the loop")
	}
}
