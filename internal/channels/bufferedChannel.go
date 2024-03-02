package channels

import "fmt"

func bufferedChannel() {
	ch := make(chan int, 1)
	ch <- 77
	val := <-ch
	fmt.Println("value received from buffered channel val:", val)
}
