package wNilZero

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("\n-- Nil and Zero Values")

		var a int
		fmt.Println(a)

		var b *int
		fmt.Println(b)

		var c bool
		fmt.Println(c)

		//var d func()
		//fmt.Println(d)

		//var e string
		//fmt.Println("[%s]", e)
	}
}
