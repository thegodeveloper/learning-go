package regex

import (
	"fmt"
	"regexp"
)

func Demo(show bool) {
	if show {
		fmt.Println("-- Regular Expressions Demo")

		// Compile a regex pattern to match email address
		// hyphen must be at the end
		re := regexp.MustCompile(`[a-zA-Z0-9._+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

		// Test strings
		email1 := "user@example.com"
		email2 := "invalid_email"
		email3 := "firstname_lastname123@example.com.co"

		// Match Emails
		fmt.Println("email1:", re.MatchString(email1))
		fmt.Println("email2:", re.MatchString(email2))
		fmt.Println("email3:", re.MatchString(email3))

		// Capturing Groups
		// Compile a regex pattern to capture date components
		re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

		// Test String
		date := "2026-06-21"

		// Match Dates
		submatches := re.FindStringSubmatch(date)
		fmt.Println(submatches)
		fmt.Println(submatches[0])
		fmt.Println(submatches[1])
		fmt.Println(submatches[2])
		fmt.Println(submatches[3])

		// Replace Characters

		// Source String
		str := "Hello World"

		re = regexp.MustCompile(`[aeiou]`)

		result := re.ReplaceAllString(str, "*")
		fmt.Printf("Source String: %s, was replaced with: %s\n", str, result)

		// Flags & Options
		// i - case insensitive
		// m - multi line model
		// s - dot matches all

		re = regexp.MustCompile(`(?i)go`)

		// Text String
		text := "Go is getting great"

		// Match
		fmt.Println("Match:", re.MatchString(text))
	}
}
