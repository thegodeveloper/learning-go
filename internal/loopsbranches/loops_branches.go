package loopsbranches

import (
	"fmt"
	"math/rand"
	"time"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Loop branches in Go")

		rand.Seed(time.Now().UnixNano())
		x := rand.Float32()

		if x < 0.5 {
			fmt.Println("head")
		} else {
			fmt.Println("tail")
		}
	}
}
