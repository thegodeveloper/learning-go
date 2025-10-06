package errors

import (
	"errors"
	"log"
	"time"
)

var (
	ErrNetwork = errors.New("network error")
	ErrInput   = errors.New("input error")
)

func namedErrors(show bool) {
	if show {
		for {
			err := dbConn()
			if err == nil {
				// success - exit the loop
				break
			}
			if errors.Is(err, ErrNetwork) {
				log.Println("recoverable network error")
				time.Sleep(1 * time.Second)
				continue
			}
		}
	}
}

func dbConn() error {
	return ErrNetwork
}
