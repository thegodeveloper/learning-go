package channels

import "fmt"

func f1(c chan int, x int) {
	fmt.Println("value of x:", x)
	c <- x
}

func f2(c chan<- int, x int) {
	fmt.Println("value of x:", x)
	c <- x
}

// ChannelsAsFunctionParameters demonstrates using channels as function parameters
func ChannelsAsFunctionParameters(show bool) {
	if !show {
		return
	}
	channelsAsFunctionParameters()
}

func channelsAsFunctionParameters() {
	c1 := make(chan int)
	c2 := make(chan int)

	go f1(c1, 1)
	go f2(c2, 2)

	x := <-c1
	y := <-c2

	fmt.Println("value received", x, y)
}
