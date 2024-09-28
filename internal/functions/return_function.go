package functions

import "fmt"

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func mainReturnFunction(show bool) {
	if show {
		fmt.Println("--- Return Function ---")

		twoBase := makeMult(2)
		threeBase := makeMult(3)
		for i := 0; i < 3; i++ {
			fmt.Println(twoBase(i), threeBase(i))
		}
	}
}
