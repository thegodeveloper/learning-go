package functions

import (
	"fmt"
	"sort"
)

// FunctionsAsParameters demonstrates passing functions as parameters
func FunctionsAsParameters(show bool) {
	if show {
		fmt.Println("--- Functions as Parameters")

		people := []Person{
			{"Pat", "Patterson", 37},
			{"Tracy", "Boddaughter", 23},
			{"Fred", "Fredson", 18},
		}

		// sort by last name
		sort.Slice(people, func(i, j int) bool {
			return people[i].LastName < people[j].LastName
		})
		fmt.Println(people)

		// sort by age
		sort.Slice(people, func(i, j int) bool {
			return people[i].Age < people[j].Age
		})
		fmt.Println(people)
	}
}
