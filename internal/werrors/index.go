package werrors

import (
	"fmt"
)

func Index(show bool) {
	if show {
		fmt.Println("--- Working with Errors ---")

		fmt.Println("-- Returning errors")
		basics()

		fmt.Println("-- Simple Errors")
		simpleErrors()

		fmt.Println("-- Sentinel Errors")
		sentinelErrors()

		fmt.Println("-- Errors are Values")
		UseLoginAndGetData()

		fmt.Println("-- Wrapping Errors")
		wrappingErrors()

		fmt.Println("-- Errors Is")
		errorsIs()
	}
}
