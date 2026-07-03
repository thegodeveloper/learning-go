package rand

import (
	"fmt"
	"math/rand"
)

func RandomNumber(show bool) {
	if show {
		fmt.Println("-- Random number --")

		// Random number between 0 and 100
		fmt.Println(rand.Intn(101))

		// Random number between 1 and 100
		fmt.Println(rand.Intn(100) + 1)
	}
}
