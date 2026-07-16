package strings

import (
	"fmt"
	"strings"
)

func StringOperations(show bool) {
	if show {
		fmt.Println("--- String Operations ---")

		var builder strings.Builder
		builder.Grow(128)
		// Write plain strings
		builder.WriteString("Go is ")
		builder.WriteString("fast!\n")
		// Append formatted data
		language := "Go"
		version := 1.26
		_, err := fmt.Fprintf(&builder, "Language: %s - ", language)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		_, err = fmt.Fprintf(&builder, "Version: %.2f - ", version)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		cores := 12
		_, err = fmt.Fprintf(&builder, "Running on %d logical CPU cores\n", cores)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		final := builder.String()
		fmt.Println("Final output: ")
		fmt.Print(final)

		// Report internal statistics
		fmt.Printf("Length: %d bytes, Capacity: %d bytes\n", builder.Len(), builder.Cap())

		// Original string
		s := "Go is powerful!"
		// Count occurrences
		fmt.Println("Count of 's':", strings.Count(s, "s"))
		// Convert to uppercase and lowercase
		fmt.Println("Upper:", strings.ToUpper(s))
		fmt.Println("Lower:", strings.ToLower(s))
		// Replace words
		replaced := strings.ReplaceAll(s, "powerful", "fast")
		fmt.Println("Replaced:", replaced)

		// Split into words
		words := strings.Fields(s)
		fmt.Println("Words:", words)
		// Join words with dashes
		dashes := strings.Join(words, "-")
		fmt.Println("Dashes:", dashes)
		// Index of a substrings
		index := strings.Index(s, "powerful")
		fmt.Println("Index of 'powerful':", index)
	}
}
