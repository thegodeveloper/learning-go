package enums

import "fmt"

type DayOfTheWeek uint8

const (
	Monday DayOfTheWeek = iota
	Tuesday
	Wednesday
	Thurday
	Friday
	Saturday
	Sunday
)

func Master(show bool) {
	if show {
		fmt.Printf("Monday is %d\n", Monday)
		fmt.Printf("Wednesday is %d\n", Wednesday)
		fmt.Printf("Friday is %d\n", Friday)
	}
}
