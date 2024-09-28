package functions

import (
	"fmt"
	"io"
	"log"
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

func printFileContent(fileName string) {
	f, closer, err := getFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func mainPrintFileContent(show bool) {
	if show {
		fmt.Println("--- print file content ---")

		printFileContent("./books.txt")
	}
}
