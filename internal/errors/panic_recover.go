package errors

import (
	"errors"
	"fmt"
	"os"
)

func PanicRecover(show bool) {
	if show {
		fmt.Println("--- Panic Recover ---")

		// Simulate file handling and check against known sentinel errors
		_, err := os.Open("/tmp/nonexisting.txt")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file does not exist")
		}

		fmt.Println("--- Panic Recover ---")
		// Attempt a risky file operation
		riskyFileOperation("/tmp/nonexisting.txt")
		fmt.Println("Program continues normally after recover.")
	}

}

func riskyFileOperation(filename string) {
	defer func() {
		// Recover from panic if it occurs
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Try to open a file that may not exist
	f, err := os.Open(filename)
	if err != nil {
		// Wrapping the error with additional context
		panic(fmt.Errorf("critical failure opening %q: %w", filename, err))
	}
	defer f.Close()

	fmt.Println("Successfully opened file:", filename)
}
