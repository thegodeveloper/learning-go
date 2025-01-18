package which

import (
	"fmt"
	"os"
	"path/filepath"
)

func Master(show bool) {
	if show {
		arguments := os.Args
		if len(arguments) == 1 {
			fmt.Println("Usage: go run ./cmd/cli/main.go command")
			return
		}

		file := arguments[1]
		path := os.Getenv("PATH")
		pathSplit := filepath.SplitList(path)

		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			fileInfo, err := os.Stat(fullPath)
			if err != nil {
				continue
			}
			mode := fileInfo.Mode()
			if !mode.IsRegular() {
				continue
			}
			if mode&0111 != 0 {
				fmt.Println(fullPath)
				return
			}
		}
	}
}
