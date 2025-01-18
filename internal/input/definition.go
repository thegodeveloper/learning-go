package input

import (
	"fmt"
	"os"
	"strconv"
)

func Master(show bool) {
	if show {
		fmt.Println("Input Definition")

		readingFromStandardInput(false)
		readingArgsFromCommandLine(false)
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

func readingArgsFromCommandLine(show bool) {
	if show {
		fmt.Println("Reading args from command line")

		arguments := os.Args
		if len(arguments) == 1 {
			fmt.Println("Need one or more arguments!")
			return
		}

		var minv, maxv float64
		var initialized int

		for i := 1; i < len(arguments); i++ {
			n, err := strconv.ParseFloat(arguments[i], 64)
			if err != nil {
				continue
			}
			if initialized == 0 {
				minv = n
				maxv = n
				initialized = 1
				continue
			}
			if n < minv {
				minv = n
			}
			if n > maxv {
				maxv = n
			}
		}

		fmt.Println("Min:", minv)
		fmt.Println("Max:", maxv)
	}
}
