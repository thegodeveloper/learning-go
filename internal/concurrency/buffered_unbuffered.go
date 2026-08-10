package concurrency

import (
	"fmt"
	"time"
)

// sendNumbers send a sequence of numbers to the provided channel.
func sendNumbers(ch chan int, count int) {
	for i := range count {
		fmt.Printf("Sending %d...\n", i)
		ch <- i
		fmt.Printf("Sent %d!\n", i)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("Closing channel")
	close(ch)
}

// receiveNumbers reads numbers from the channel until it is closed.
func receiveNumbers(ch chan int) {
	for num := range ch {
		fmt.Printf("Received %d\n", num)
		time.Sleep(time.Millisecond * 400)
	}
	fmt.Println("Receiver done")
}

func BufferedUnbuffered(show bool) {
	if show {
		fmt.Println("--- Buffered unbuffered ---")

		fmt.Println("--- Unbuffered Channel Example ---")
		unbuffered := make(chan int)
		go sendNumbers(unbuffered, 3)
		receiveNumbers(unbuffered) // Blocks until all values are consumed.

		fmt.Println("--- Buffered Channel Example ---")
		buffered := make(chan int, 3) // Capacity of 3

		go func() {
			sendNumbers(buffered, 5)
		}()
		receiveNumbers(buffered)
	}
}
