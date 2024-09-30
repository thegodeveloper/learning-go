package defersimplecat

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Master(show bool) {
	if show {
		fmt.Println("--- Defer Simple Cat")

		if len(os.Args) < 2 {
			log.Fatal("no file specified")
		}
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)
		data := make([]byte, 2048)
		for {
			count, err := f.Read(data)
			os.Stdout.Write(data[:count])
			if err != nil {
				if err == io.EOF {
					log.Fatal(err)
				}
				break
			}
		}
	}
}
