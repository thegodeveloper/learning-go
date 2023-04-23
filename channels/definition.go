package channels

import "fmt"

func definition() {
	fmt.Println("-- Buffered Channel")
	bufferedChannel()

	fmt.Println("-- Unbuffered Channel")
	unbufferedChannel()

	fmt.Println("-- Read and Write from/to the Channel")
	readAndWriteToFromChannel()

	fmt.Println("-- Closing a Channel")
	closedChannel()
}

func bufferedChannel() {
	ch := make(chan int, 1)
	ch <- 77
	val := <-ch
	fmt.Println("value received from buffered channel val:", val)
}

func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		ch <- 77
	}()
	val := <-ch
	fmt.Println("value received from unbuffered channel val:", val)
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

func closedChannel() {
	ch := make(chan int, 7)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	ch <- 7
	close(ch)

	for {
		val, ok := <-ch
		if !ok { // if the channel is closed, ok will be false
			fmt.Println("channel closed!")
			break
		}
		fmt.Println("receiving from channel:", val)
	}
}
