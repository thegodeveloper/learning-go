package randommessage

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, world", "Καλημέρα κόσμε", "سلام دنیا", "Привет, мир",
}

func Index(show bool) {
	if show {
		rand.NewSource(time.Now().UnixNano())
		index := rand.Intn(len(helloList))
		msg, err := hello(index)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range:" + strconv.Itoa(index))
	}
	return helloList[index], nil
}
