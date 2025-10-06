package passingarguments

import (
	"fmt"
	"os"
)

func Run(show bool) {
	if show {
		argsWithProg := os.Args
		argsWithoutProg := os.Args[1:]
		arg := os.Args[3]

		fmt.Println(argsWithProg)
		fmt.Println(argsWithoutProg)
		fmt.Println(arg)
	}
}
