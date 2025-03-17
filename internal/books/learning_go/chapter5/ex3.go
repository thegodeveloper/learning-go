package chapter5

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(body string) string {
		return prefix + " " + body
	}
}

func ex3(show bool) {
	if show {
		helloPrefix := prefixer("Hello")
		fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
		fmt.Println(helloPrefix("Maria")) // should print Hello Maria
	}
}
