package customtype

import "fmt"

type CarModel string

func Run(show bool) {
	if show {
		fmt.Println("-- Custom Type")

		myCar := CarModel("Chevelle")
		fmt.Printf("the value of myCar is: %v\n", myCar)
	}
}
