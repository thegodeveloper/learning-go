package goroutines

import (
	"fmt"
	"net/http"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Goroutines")

		implementCountTo()
		implementCancelFunctionTerminateGoroutine()
		implementBackpressure()
	}
}

func implementCountTo() {
	for i := range countTo(7) {
		fmt.Println(i)
	}
}

func implementCancelFunctionTerminateGoroutine() {
	fmt.Println("-- Implement Cancel Function Terminate Goroutine")
	ch, cancel := countTo2(7)
	for i := range ch {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}

// Don't do this in a real program
// This violates one of our "when to use concurrency" guidelines
func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func implementBackpressure() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
