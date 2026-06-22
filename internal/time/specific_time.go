package time

import (
	"fmt"
	"time"
)

func SpecificTime(show bool) {
	if show {
		fmt.Println("-- Specific Time")

		// Current loca time
		fmt.Println(time.Now())

		// Specific Time
		specificTime := time.Date(2026, time.June, 22, 7, 7, 7, 7, time.UTC)
		fmt.Println("Specific Time:", specificTime)

		// Parse Time
		// Go reference time: Mon Jan 2 15:04:05 MST 2006
		// layout 2006-01-02
		parsedTime, err := time.Parse("2006-01-02", "2020-05-01")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Parse Time Layout 2006-01-02:", parsedTime)

		// layout 06-01-02
		parsedTime, err = time.Parse("06-01-02", "20-05-01")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Parse Time Layout 06-01-02:", parsedTime)

		// layout 06-1-2 15-04
		parsedTime, err = time.Parse("06-1-2 15-04", "20-5-1 18-03")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Parse Time Layout 06-1-2 15-04:", parsedTime)

		// Formatting Time
		t := time.Now()
		fmt.Println("Formatted time:", t.Format("Monday 2006-01-02 15:04:05"))
	}
}
