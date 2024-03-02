package channels

import "fmt"

func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		ch <- 77
	}()
	val := <-ch
	fmt.Println("value received from unbuffered channel val:", val)
}
