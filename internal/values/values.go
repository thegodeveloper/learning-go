package values

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("go" + "lang")

		fmt.Println("1+1 = ", 1+1)
		fmt.Println("7.0/3.0 = ", 7.0/3.0)

		fmt.Println(true && false)
		fmt.Println(true || false)
		fmt.Println(!true)
	}
}
