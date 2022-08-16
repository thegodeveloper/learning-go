package main

import "fmt"

func something() {
	defer fmt.Println("closed something")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i > 2 {
			panic("panic was called")
		}
	}
}

func main() {
	defer fmt.Println("close main")
	something()
	fmt.Println("something was finished") // this is not called
}

