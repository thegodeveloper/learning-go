package functions

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
