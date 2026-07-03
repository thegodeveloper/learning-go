package time

import (
	"fmt"
	"time"
)

func Formatting(show bool) {
	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02T15:04:05Z07:00"
	str := "2026-07-02T22:39:00Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time", err)
		return
	}
	fmt.Println(t.Format(layout))

	str1 := "Jul 02, 2026 10:42:00 PM"
	layout1 := "Jan 02, 2006 03:04:05 PM"
	t1, err := time.Parse(layout1, str1)
	if err != nil {
		fmt.Println("Error parsing time", err)
	}
	fmt.Println(t1.Format(layout1))
}
