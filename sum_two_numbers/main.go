package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithProg := os.Args

	numa, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	numb, err := strconv.Atoi(argsWithProg[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	res := numa + numb
	fmt.Printf("numa: %d + numb: %d = %d\n", numa, numb, res)
}
