package channels

import "fmt"

// ClosedChannel demonstrates closing channels and detecting closed state
func ClosedChannel(show bool) {
	if !show {
		return
	}
	closedChannel()
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
