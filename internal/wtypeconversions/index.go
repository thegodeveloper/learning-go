package wtypeconversions

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("-- Type Conversions")

		definition(false)
		integerTypeConversion(false)
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

func integerTypeConversion(show bool) {
	if show {
		var x int = 10
		var b byte = 100
		var sum3 int = x + int(b)
		var sum4 byte = byte(x) + b
		fmt.Println("sum3 = ", sum3)
		fmt.Println("sum4 = ", sum4)
	}
}
