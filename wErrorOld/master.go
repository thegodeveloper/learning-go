package wErrorOld

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var Musketeers = []string{
	"Athos", "Porthos", "Aramis", "D'Artagnan",
}

func GetMusketeer(id int) (string, error) {
	if id < 0 || id >= len(Musketeers) {
		return "", errors.New(fmt.Sprintf("Invalid id [%d]", id))
	}
	return Musketeers[id], nil
}

func demo(show bool) {
	if show {
		fmt.Println("\n-- Error Demo")

		rand.Seed(time.Now().UnixNano())
		id := rand.Int() % 6

		mosq, err := GetMusketeer(id)
		if err == nil {
			fmt.Printf("[%d] %s", id, mosq)
		} else {
			fmt.Println(err)
		}

		fmt.Println("Hello")
	}
}

func Master(show bool) {
	if show {
		fmt.Println("\n-- Errors")

		demo(false)
		errorsAndValues(false)
		errorFromFunction(false)
		wrappingErrors(false)
		simple(false)
		sentinelError(false)
	}
}
