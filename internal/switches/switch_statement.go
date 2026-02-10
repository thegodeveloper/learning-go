package switches

import "fmt"

func SwitchStatement(show bool) {
	if show {
		fmt.Println("--- Switch Statement ---")

		words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

		for _, word := range words {
			switch size := len(word); size {
			case 1, 2, 3, 4:
				fmt.Println(word, "is a short word!")
			case 5:
				wordLen := len(word)
				fmt.Println(word, "is exactly the right length:", wordLen)
			case 6, 7, 8, 9:
			default:
				fmt.Println(word, "is a long word!")
			}
		}

		switchCase(true)
		multipleSwitchCases(true)
	}
}

func switchCase(show bool) {
	if show {
		fruit := "pineapple"

		switch fruit {
		case "apple":
			fmt.Println("It's an apple!")
		case "banana":
			fmt.Println("It's a banana!")
		default:
			fmt.Println("Unknown fruit!")
		}
	}
}

func multipleSwitchCases(show bool) {
	if show {
		day := "Monday"

		switch day {
		case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
			fmt.Println("It's a weekday!")
		case "Saturday", "Sunday":
			fmt.Println("It's a weekend day!")
		default:
			fmt.Println("Unknown day!")
		}
	}
}
