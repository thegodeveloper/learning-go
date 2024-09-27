package switches

import (
	"fmt"
	"math/rand"
	"time"
)

func switchWithConditions(show bool) {
	if show {
		fmt.Println("--- Switch Conditions ---")

		rand.Seed(time.Now().UnixNano())
		x := rand.Float32()

		switch {
		case x < 0.25:
			fmt.Println("Q1")
		case x < 0.5:
			fmt.Println("Q2")
		case x < 0.75:
			fmt.Println("Q3")
		default:
			fmt.Println("Q4")
		}
	}
}
