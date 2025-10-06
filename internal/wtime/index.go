package wtime

import (
	"fmt"
	"log"
	"time"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Time")

		t, err := time.Parse("2006-01-02 15:04:05 -0700", "2017-17-07 00:00:00 +0000")
		if err != nil {
			log.Fatalf("error parsing the date: %s\n", err)
		}
		fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	}
}
