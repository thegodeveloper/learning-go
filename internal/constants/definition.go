package constants

import "fmt"

const monthsInYear = 12
const str = "hello world"
const num = 3
const num64 int64 = 3

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
	}
}
