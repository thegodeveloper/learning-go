package wgoroutines

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Master(show bool) {
	if show {
		fmt.Println("-- Goroutines")

		evenCh, oddCh := make(chan bool, 1), make(chan bool, 1)
		defer close(evenCh)
		defer close(oddCh)

		wg = sync.WaitGroup{}
		wg.Add(2)

		go even(evenCh, oddCh)
		go odd(oddCh, evenCh)

		evenCh <- true
		wg.Wait()
	}
}

func even(evenCh, oddch chan bool) {
	for i := 2; i <= 10; i += 2 {
		<-oddch
		fmt.Println(i)
		evenCh <- true
	}

	wg.Done()
}

func odd(oddCh, evenCh chan bool) {
	for i := 1; i <= 10; i += 2 {
		<-evenCh
		fmt.Println(i)
		oddCh <- true
	}

	wg.Done()
}
