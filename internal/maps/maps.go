package maps

import "fmt"

func Master(show bool) {
	if show {
		definition(false)

		mapsDeclarations(false)

		commaOkIdiomInMaps(false)

		delitingFromMaps(false)

		mapReadWrite(false)
	}
}

func definition(show bool) {
	if show {
		fmt.Println("--- Definition")

		// declaring a map using make:
		// create a map with string keys and stores data that is an int type
		// 10 signifies that we want to pre-size for 10 entries
		var counters = make(map[string]int, 10)
		fmt.Printf("counters length: %d\n", len(counters))
		counters["one"] = 1
		counters["two"] = 2
		fmt.Printf("counters length: %d\n", len(counters))

		// declaring a map using composite literal:
		modelToMake := map[string]string{
			"prius":    "toyota",
			"chevelle": "chevy",
		}
		fmt.Printf("modelToMake length: %d\n", len(modelToMake))

		// accessing values:
		carMake := modelToMake["chevelle"]
		fmt.Printf("car maker: %s\n", carMake)

		// non-existing key returns an empty string
		carMake = modelToMake["non-existing"]
		fmt.Printf("value of non-existing key: %s\n", carMake)

		// other way to validate if the value exists
		if carMake, ok := modelToMake["non-existing"]; ok {
			fmt.Printf("car model \"non-existing\" has make %q\n", carMake)
		} else {
			fmt.Printf("car model \"non-existing\" has an unknown make\n")
		}

		// adding values
		modelToMake["outback"] = "subaru"
		counters["pagehits"] = 10

		fmt.Printf("counters length: %d\n", len(counters))
		fmt.Printf("modelToMake length: %d\n", len(modelToMake))

		// print all the values of modelToMake
		for key, val := range modelToMake {
			fmt.Printf("car model %q has make %q\n", key, val)
		}
	}
}

func mapsDeclarations(show bool) {
	if show {
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
