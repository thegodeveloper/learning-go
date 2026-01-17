package closures

import "fmt"

// Declaration demonstrates closure definition and usage
func Declaration(show bool) {
	if show {
		fmt.Println("--- Closure Definition")

		a := 20
		f := func() {
			fmt.Println(a)
			a = 30
		}
		f()
		fmt.Println(a)
	}
}
