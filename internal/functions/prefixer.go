package functions

import "fmt"

func prefixer(p string) func(string) string {
	return func(s string) string {
		return p + " " + s
	}
}

func mainPrefixer(show bool) {
	if show {
		helloPrefix := prefixer("hello")
		fmt.Println(helloPrefix("Box"))
		fmt.Println(helloPrefix("Maria"))
	}
}
