package unicode

import (
	"fmt"
	"unicode/utf8"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Unicode in Go")

		s := "Hello, 世界"
		fmt.Println(s)

		pf := HasPrefix(s, "Hello,")
		if pf {
			fmt.Println("Prefix Contained")
		} else {
			fmt.Println("Prefix not contained")
		}

		sf := HasSuffix(s, "世界")
		if sf {
			fmt.Println("Suffix Contained")
		} else {
			fmt.Println("Suffix not contained")
		}

		c := Contains(s, "llo")
		if c {
			fmt.Println("Value contained")
		} else {
			fmt.Println("Value not contained")
		}

		for i := 0; i < len(s); {
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%d\t%c\n", i, r)
			i += size
		}

		for i, r := range s {
			fmt.Printf("%d\t%qt%d\n", i, r, r)
		}

		fmt.Printf("number of runes in a string: %d\n", utf8.RuneCountInString(s))
	}
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
