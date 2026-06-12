package maps

import "fmt"

func MapNil(show bool) {
	if show {
		fmt.Println("Map Nil Initialized")

		var objectDesc map[string]string

		if objectDesc == nil {
			fmt.Println("The map is initialized to nil value.")
		} else {
			fmt.Println("The map is not initialized to nil value.")
		}

		// query a nil map
		val := objectDesc["none"]
		fmt.Println("Attempting to access a map nil value:", val)

		// this produces an error
		// objectDesc["macbook"] = "The MacBook Pro Max is the best"
		// fmt.Println("values on objectDesc map:", objectDesc)

		// To be able to add values to the map we need to initialize it
		objectDesc = make(map[string]string)
		objectDesc["macbook"] = "The MacBook Pro Max is the best"
		fmt.Println("values on objectDesc map:", objectDesc)
	}
}
