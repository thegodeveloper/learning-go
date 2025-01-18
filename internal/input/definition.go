package input

import "fmt"

func Master(show bool) {
	if show {
		fmt.Println("Input Definition")

		readingFromStandardInput(false)
	}
}

func readingFromStandardInput(show bool) {
	if show {
		fmt.Println("Reading from standard input")

		fmt.Printf("Please give me your name:")
		var name string
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Error reading input")
		} else {
			fmt.Println("Your name is:", name)
		}
	}
}
