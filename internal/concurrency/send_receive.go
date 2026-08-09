package concurrency

import (
	"fmt"
	"time"
)

func sender(ch chan int, count int) {
	for i := range count {
		fmt.Printf("Sending: sending %d\n", i)
		ch <- i // send data to the channel
		time.Sleep(time.Millisecond * 300)
	}
	fmt.Println("Sender: closing channel")
	close(ch) // signal that no more data will be sent.
}

// receiver reads numbers from the channel until it is closed.
func receiver(ch chan int, done chan bool) {
	for num := range ch { // Automatically stops when channel is closed.
		fmt.Printf("Receiver: received %d\n", num)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("Receiver: no more data, stopping")
	done <- true
}

func SendReceive(show bool) {
	if show {
		fmt.Println("--- Concurrency SendReceive ---")

		ch := make(chan int)    // Unbuffered channel for exchanging data.
		done := make(chan bool) // Used to signal completion.

		// Start sender and receiver goroutines.
		go sender(ch, 5)
		go receiver(ch, done)

		// Wait for the receiver to finish.
		<-done
		fmt.Println("--- All done! ---")
	}
}
