package main

import "fmt"

func main() {
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
