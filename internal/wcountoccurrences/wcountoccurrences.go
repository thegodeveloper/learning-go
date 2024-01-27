package wcountoccurrences

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Master(show bool) {
	if show {
		// Open the file
		file, err := os.Open("wCountOccurrencesInFile/filename.csv")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		// Initialize the counter
		count := 0

		// Read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, ",")
			// Search for "william" in each field
			for _, field := range fields {
				trimmedField := strings.TrimSpace(field)
				if trimmedField == "william" {
					count++
					break
				}
			}
		}

		// Print the total number of occurrences
		fmt.Println("Total number of items found:", count)
	}
}
