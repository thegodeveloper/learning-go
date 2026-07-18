package regex

import (
	"fmt"
	"regexp"
)

func RegularExpressions(show bool) {
	if show {
		fmt.Println("--- Regular Expressions ---")

		// Simple matching using MatchString
		fmt.Println("--- Simple Matching ---")
		pattern := `Go`
		re := regexp.MustCompile(pattern)
		text := "Go is expressive, concise, clean, and efficient."
		fmt.Printf("Does the text contains '%s'? %v\n", pattern, re.MatchString(text))

		// Finding all matches
		fmt.Println("--- Finding All Matches ---")
		pattern = `[A-Z][a-z]+`
		re = regexp.MustCompile(pattern)
		text = "Go Rust Zig C C++"
		matches := re.FindAllString(text, -1)
		fmt.Printf("Capitalized words: %v\n", matches)

		// Extracting submatches (groups)
		fmt.Println("--- Extracting Submatches ---")
		pattern = `(\w+)@(\w+\.\w+)`
		re = regexp.MustCompile(pattern)
		email := "contact us at support@example.com"
		submatches := re.FindStringSubmatch(email)
		if len(submatches) > 0 {
			fmt.Printf("Full match: %s\n", submatches[0])
			fmt.Printf("Username: %s\n", submatches[1])
			fmt.Printf("Domain: %s\n", submatches[2])
		}

		// Replacing Matches
		fmt.Println("--- Replacing Matches ---")
		pattern = `\d+`
		re = regexp.MustCompile(pattern)
		text = "Order IDs: 1234, 5678, 91011"
		replaced := re.ReplaceAllString(text, "XXXX")
		fmt.Println("Replaced IDs:", replaced)

		// Validating Input
		fmt.Println("--- Validating Input ---")
		pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re = regexp.MustCompile(pattern)
		emails := []string{"user@example.com", "invalid-email@", "hello.world@ziglang.org"}
		for _, e := range emails {
			if re.MatchString(e) {
				fmt.Printf("%s is a valid email\n", e)
			} else {
				fmt.Printf("%s is not a valid email\n", e)
			}
		}
	}
}
