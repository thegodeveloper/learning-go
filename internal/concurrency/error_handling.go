package concurrency

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func ErrorHandling(show bool) {
	if show {
		fmt.Println("-- Concurrency Error-Handling")

		done := make(chan interface{})
		defer close(done)

		urls := []string{"https://example.com", "https://golang.org", "https://doesnot.exist"}

		errCount := 0
		for res := range checkStatus(done, urls...) {
			if res.Error != nil {
				fmt.Printf("error: %v\n", res.Error)
				errCount++
				if errCount >= 2 {
					fmt.Println("too many errors, exiting")
					break
				}
				continue
			}
			fmt.Printf("Status: %d\n", res.Response.StatusCode)
		}
	}
}

// checkStatus queries multiple URLs and report each Response or error
func checkStatus(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)
		for _, url := range urls {
			resp, err := http.Get(url)
			select {
			case <-done:
				return
			case results <- Result{Error: err, Response: resp}:
			}
		}
	}()
	return results
}
