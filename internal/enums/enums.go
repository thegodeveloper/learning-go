package enums

import "fmt"

type DayOfTheWeek uint8

const (
	Monday DayOfTheWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func Run(show bool) {
	if show {
		fmt.Println("-- Enums in Go")

		fmt.Printf("Monday is %d\n", Monday)
		fmt.Printf("Tuesday is %d\n", Tuesday)
		fmt.Printf("Wednesday is %d\n", Wednesday)
		fmt.Printf("Thursday is %d\n", Thursday)
		fmt.Printf("Friday is %d\n", Friday)
		fmt.Printf("Saturday is %d\n", Saturday)
		fmt.Printf("Sunday is %d\n", Sunday)
	}
}
