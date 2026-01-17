package channels

import "fmt"

// ReadAndWriteToFromChannel demonstrates reading and writing to channels
func ReadAndWriteToFromChannel(show bool) {
	if !show {
		return
	}
	readAndWriteToFromChannel()
}

func readAndWriteToFromChannel() {
	ch := make(chan int)

	// write to the channel
	go func() {
		ch <- 77
	}()

	// read from the channel
	val := <-ch
	fmt.Println("value read from the channel:", val)

	// write to the channel again
	go func() {
		ch <- 777
	}()

	// read from the channel again
	val = <-ch
	fmt.Println("value read from the channel again:", val)
}
