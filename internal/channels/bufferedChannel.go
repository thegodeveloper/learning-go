package channels

import "fmt"

// BufferedChannel demonstrates buffered channel usage
func BufferedChannel(show bool) {
	if !show {
		return
	}
	bufferedChannel()
}

func bufferedChannel() {
	ch := make(chan int, 1)
	ch <- 77
	val := <-ch
	fmt.Println("value received from buffered channel val:", val)
}
