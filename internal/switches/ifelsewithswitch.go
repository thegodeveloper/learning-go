package switches

import (
	"fmt"
	"math/rand"
)

func ifElseWithSwitch(show bool) {
	if show {
		fmt.Println("--- Rewrite an if/else with a switch ---")

		switch n := rand.Intn(10); {
		case n == 0:
			fmt.Println("That's too low")
		case n > 5:
			fmt.Println("That's too big:")
		default:
			fmt.Println("That's a good number:", n)
		}
	}
}
