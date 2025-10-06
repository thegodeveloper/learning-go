package openfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(show bool) {
	if show {
		// open the file for reading
		file, err := os.Open("file.txt")
		if err != nil {
			fmt.Println("Error openning file:", err)
			return
		}
		defer file.Close()

		// create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// split each line into fields separated by commas
			fields := strings.Split(scanner.Text(), ",")

			// print each field
			for _, field := range fields {
				fmt.Println(field)
			}
		}

		// check for any errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning file:", err)
		}
	}
}
