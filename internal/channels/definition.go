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
