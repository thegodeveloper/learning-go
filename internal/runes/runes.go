package runes

import "fmt"

func Run(show bool) {
	if show {
		// declaring and initializing a Unicode character
		emoji := '😀'

		// %T represents the type of the variable num
		fmt.Printf("Data type of %c is %T and the rune value is %U \n", emoji, emoji, emoji)
	}
}
