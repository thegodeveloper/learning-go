package main

import "fmt"

var x = xSetter()

func xSetter() int {
	fmt.Println("xSetter")
	return 77
}

func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println("this is the main function")
}
