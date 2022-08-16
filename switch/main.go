package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var finger int = 1

	switch finger {
	case 0:
		fmt.Println("Thumb")
	case 1:
		fmt.Println("Index")
	case 2:
		fmt.Println("Middle")
	case 3:
		fmt.Println("Ring")
	case 4:
		fmt.Println("Pinkie")
	default:
		fmt.Println("Humans usually have no more than five fingers")
	}

	switchWithConditions()
}

func switchWithConditions() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Float32()

	switch {
	case x < 0.25:
		fmt.Println("Q1")
	case x < 0.5:
		fmt.Println("Q2")
	case x < 0.75:
		fmt.Println("Q3")
	default:
		fmt.Println("Q4")
	}
}
