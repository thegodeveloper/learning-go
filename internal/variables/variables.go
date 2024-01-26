package variables

import "fmt"

func Master(show bool) {
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

		variables()
	}
}

func variables() {
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
