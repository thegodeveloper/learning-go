package strings

import "fmt"

func ToBeAwareInStrings(show bool) {
	if show {
		fmt.Println("--- to be aware of working with strings ---")
		s := "Hello ★"
		fmt.Println("s:", s)
		fmt.Println("s lenght:", len(s))
		fmt.Println("the ★ occupies 3 bytes")
	}
}
