package wcustomtype

import "fmt"

type CarModel string

func Master(show bool) {
	if show {
		fmt.Println("\n-- Custom Type")

		myCar := CarModel("Chevelle")
		fmt.Printf("the value of myCar is: %v\n", myCar)
	}
}
