package switches

import "fmt"

func BasicBlankSwitch(show bool) {
	if show {
		words := []string{"hi", "salutations", "seven", "hello"}

		for _, word := range words {
			switch wordLen := len(word); {
			case wordLen < 5:
				fmt.Println(word, "is a short word!")
			case wordLen > 10:
				fmt.Println(word, "is a long word!")
			default:
				fmt.Println(word, "is exactly the right length!")
			}
		}
	}
}
