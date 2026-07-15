package strings

import (
	"fmt"
	"unicode/utf8"
)

func StringsRunes(show bool) {
	if show {
		fmt.Println("--- StringRunes ---")

		str := "Go ♥ Zig"
		fmt.Println("Original string:", str)
		fmt.Println("Length in bytes:", len(str))

		fmt.Println("*** Byte-wise iteration:")
		for i, r := range str {
			fmt.Printf("Rune at byte %d: %c (U+%04X)\n", i, r, r)
		}

		// Convert to []byte and []rune
		bytes := []byte(str)
		runes := []rune(str)

		fmt.Println("Converted to []byte:", bytes)
		fmt.Println("Converted to []rune:", runes)

		fmt.Println("Number of runes:", utf8.RuneCountInString(str))
		fmt.Println("First rune:", string(runes[0]), "| Uni:", runes[0])
	}
}
