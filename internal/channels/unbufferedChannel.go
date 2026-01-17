package channels

import "fmt"

// UnbufferedChannel demonstrates unbuffered channel usage
func UnbufferedChannel(show bool) {
	if !show {
		return
	}
	unbufferedChannel()
}

func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		ch <- 77
	}()
	val := <-ch
	fmt.Println("value received from unbuffered channel val:", val)
}
