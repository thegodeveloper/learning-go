package constants

import (
	"fmt"
	"math"
)

const monthsInYear = 12
const str = "hello world"
const num = 3
const num64 int64 = 3
const s string = "constant"

type specialStr string

func printSpecial(str specialStr) {
	fmt.Println(string(str))
}

func Master(show bool) {
	if show {
		println("Constants")
		println("Constant value: ", monthsInYear)

		const constHelloWorld = "hello world"
		printSpecial(constHelloWorld)
		printSpecial("hello world")

		fmt.Println("s value:", s)

		const n = 500_000_000
		const d = 3e20 / n
		fmt.Println("d value:", d)
		fmt.Println(int64(d))
		fmt.Println(math.Sin(d))
	}
}
