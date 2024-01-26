package strings

import "fmt"

func Master(show bool) {
	if show {
		extractingSingleValue()
		toBeAwareInStrings()
	}
}

func extractingSingleValue() {
	fmt.Println("--- extracting a single value from a string ---")
	s := "Hello Golang!"
	var b byte = s[7]
	fmt.Println("s:", s)
	fmt.Println("b:", b)

	s2 := s[4:7]
	s3 := s[:5]
	s4 := s[6:]
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
}

func toBeAwareInStrings() {
	fmt.Println("--- to be aware of working with strings ---")
	s := "Hello ★"
	fmt.Println("s:", s)
	fmt.Println("s lenght:", len(s))
	fmt.Println("the ★ occupies 3 bytes")
}
