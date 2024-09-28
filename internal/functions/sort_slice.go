package functions

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func mainSortSlice(show bool) {
	if show {
		people := []Person{
			{"Pat", "Patterson", 37},
			{"Tracy", "Bobbert", 27},
			{"Fred", "Fredson", 47},
		}
		fmt.Println("--- current people value ---")
		fmt.Println(people)

		// sort by last name
		fmt.Println("--- sorted by last name ---")
		sort.Slice(people, func(i int, j int) bool {
			return people[i].LastName < people[j].LastName
		})
		fmt.Println(people)

		// sort by age
		fmt.Println("--- sorted by age ---")
		sort.Slice(people, func(i int, j int) bool {
			return people[i].Age < people[j].Age
		})
		fmt.Println(people)
	}
}
