package wConcurrency

import (
	"errors"
	"fmt"
	"time"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Concurrency")
		implementation()
	}
}

func implementation() {
	result, err := timeLimit()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() (int, error) {
	time.Sleep(3 * time.Second)
	return 7, nil
}
