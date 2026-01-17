package maps

import "fmt"

func MapsDeclarations(show bool) {
	if show {
		fmt.Println("--- Maps Declarations")

		// nilMap is declared to be a map with string keys and int values
		// the zero value for a map is nil
		// a nil map has a length of 0
		// attempting to read a nil map always returns the zero value for the map's value type
		fmt.Println("--- nil map ---")
		var nilMap map[string]int
		fmt.Println("nilMap:", nilMap)

		// empty map literal
		// is not the same as a nil map
		// it has a length of zero
		// you can read and write to a map assigned an empty map literal
		fmt.Println("--- map literal ---")
		totalWins := map[string]int{}
		fmt.Println("totalWins:", totalWins)

		fmt.Println("--- nonempty map literal ---")
		teams := map[string][]string{
			"Orcas":   {"Fred", "Ralph", "Bijou"},
			"Lions":   {"Sarah", "Peter", "Billie"},
			"Kittens": {"Waldo", "Raul", "Ze"},
		}
		fmt.Println("teams:", teams)
		fmt.Println("teams.Orcas:", teams["Orcas"])
		fmt.Println("teams.Lions.Peter:", teams["Lions"][1])

		fmt.Println("--- map with a default size ---")
		ages := make(map[int][]string, 10)
		fmt.Println("ages:", ages)
		fmt.Println("ages length:", len(ages))
	}
}
