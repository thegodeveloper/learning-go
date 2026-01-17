package functions

import (
	"fmt"
	"io"
	"os"
)

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

func printFileContent(fileName string) error {
	f, closer, err := getFile(fileName)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer closer()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("error reading file: %w", err)
			}
			break
		}
	}
	return nil
}

// PrintFileContent demonstrates returning cleanup functions
func PrintFileContent(show bool) {
	if show {
		fmt.Println("--- print file content ---")

		// Try to read books.txt from the functions directory
		err := printFileContent("internal/functions/books.txt")
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("\nThis example requires a 'books.txt' file in the internal/functions directory.")
			fmt.Println("Demonstrating with a simple example instead:\n")

			// Show a simple example of the closure pattern
			fmt.Println("Example of returning a cleanup function:")
			cleanup := func() string {
				return "Cleanup function executed"
			}
			fmt.Println(cleanup())
		}
	}
}
