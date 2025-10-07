package time

import (
	"fmt"
	"time"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Time in Go")

		t, err := time.Parse("2006-01-02 15:04:05 -0700", "2017-17-07 00:00:00 +0000")
		if err != nil {
			fmt.Printf("error parsing the date: %s\n", err)
		}
		fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	}
}
