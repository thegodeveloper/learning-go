package logfatal

import (
	"log"
	"os"
)

func Master(show bool) {
	if show {
		if len(os.Args) != 1 {
			log.Fatal("Fatal: Hello World!")
		}
		log.Panic("Panic: Hello World!")
	}
}
