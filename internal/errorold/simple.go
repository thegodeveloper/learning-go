package errorold

import (
	"fmt"
	"os"
	"strconv"
)

// Simple demonstrates simple error handling with strconv
func Simple(show bool) {
	if show {
		age := os.Args[1]

		n, err := strconv.Atoi(age)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
	}
}
