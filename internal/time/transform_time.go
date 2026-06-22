package time

import (
	"fmt"
	"time"
)

func TransformTime(show bool) {
	if show {
		fmt.Println("-- Transform Time")

		// Formatting Time
		t := time.Now()
		fmt.Println("Formatted time:", t.Format("Monday 2006-01-02 15:04:05"))

		// Add One Day
		oneDayLater := t.Add(time.Hour * 24)
		fmt.Println("One day later:", oneDayLater)

		// Round it to nearest hour
		roundedTime := t.Round(time.Hour)
		fmt.Println("Rounded time:", roundedTime)

		// Time Location
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println(err)
			return
		}
		t = time.Date(2026, time.June, 22, 7, 7, 7, 00, time.UTC)

		// Convert the time to a specific time zone
		tLocal := t.In(loc)

		// Perform rounding
		roundedTime = t.Round(time.Hour)
		roundedTimeLocal := roundedTime.In(loc)
		fmt.Println("Original Time UTC:", t)
		fmt.Println("Original Time Local:", tLocal)
		fmt.Println("Rounded Time UTC:", roundedTime)
		fmt.Println("Rounded Time Local:", roundedTimeLocal)
	}
}
