package wErrors

import (
	"fmt"
)

func Master(show bool) {
	if show {
		fmt.Println("--- Working with Errors ---")

		fmt.Println("-- Returning errors")
		basics()

		fmt.Println("-- Simple Errors")
		simpleErrors()

		fmt.Println("-- Sentinel Errors")
		sentinelErrors()

		fmt.Println("-- Errors are Values")
		_, err := LoginAndGetData("bill", "xyz", "mozart")
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("-- Wrapping Errors")
		wrappingErrors()
	}
}
