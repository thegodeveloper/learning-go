package variables

import (
	"fmt"
	"math"
)

const value = 10

func Master(show bool) {
	if show {
		declaration(false)

		variables(false)

		exe01(false)

		exe02(false)

		exe03(true)
	}
}

func declaration(show bool) {
	if show {
		var a int
		a = 42

		var aa int = 100

		b := -42

		c := "this is a string"

		var d, e string
		d, e = "var d", "var e"

		f, g := true, false

		fmt.Println(a)
		fmt.Println(aa)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
		fmt.Println(e)
		fmt.Println(f)
		fmt.Println(g)
	}
}

func variables(show bool) {
	if show {
		var a = "initial"
		fmt.Println(a)

		var b, c int = 1, 2
		fmt.Println(b, c)

		var d = true
		fmt.Println(d)

		var e int
		fmt.Println(e)

		f := "apple"
		fmt.Println(f)
	}
}

func exe01(show bool) {
	if show {
		i := 20
		f := float64(i)

		fmt.Println("value of i:", i)
		fmt.Println("value of f:", f)
	}
}

func exe02(show bool) {
	if show {
		var i = value
		var f float64 = value

		fmt.Println("value of i:", i)
		fmt.Println("value of f:", f)
	}
}

func exe03(show bool) {
	if show {
		var b uint8 = math.MaxUint8
		var smallI int32 = math.MaxInt32
		var bigI uint64 = math.MaxUint64

		fmt.Println("Max value of b:", b)
		fmt.Println("Max value of smallI:", smallI)
		fmt.Println("Max value of bigI:", bigI)

		b = b + 1
		smallI = smallI + 1
		bigI = bigI + 1

		fmt.Println("b when add 1:", b)
		fmt.Println("smallI when add 1:", smallI)
		fmt.Println("bigI when add 1:", bigI)
	}
}
