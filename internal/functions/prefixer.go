package functions

import "fmt"

func prefixer(p string) func(string) string {
	return func(s string) string {
		return p + " " + s
	}
}

// Prefixer demonstrates closures with the prefixer example
func Prefixer(show bool) {
	if show {
		helloPrefix := prefixer("hello")
		fmt.Println(helloPrefix("Box"))
		fmt.Println(helloPrefix("Maria"))
	}
}
