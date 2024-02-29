package wtypeconversions

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("-- Type Conversions")

		definition(false)
	}
}

func definition(show bool) {
	if show {
		fmt.Println("--- Definition")

		var x int = 10
		var y float64 = 30.2
		var sum1 float64 = float64(x) + y
		var sum2 int = x + int(y)

		fmt.Println("sum1 = ", sum1)
		fmt.Println("sum2 = ", sum2)
	}
}
