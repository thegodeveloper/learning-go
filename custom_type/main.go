package main

import "fmt"

type CarModel string

func main() {
	myCar := CarModel("Chevelle")
	fmt.Printf("the value of myCar is: %v\n", myCar)
}
