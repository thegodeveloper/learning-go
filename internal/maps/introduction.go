package maps

import "fmt"

func Introduction(show bool) {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a map literal
	// mapVariable = map[keyType]valueType{
	//   key1: value1,
	//   key2: value2
	// }
	if show {
		fmt.Println("-- Maps Introduction")

		familyAges := make(map[string]int)
		familyAges["John"] = 51
		familyAges["Jane"] = 47
		familyAges["Mike"] = 29
		familyAges["Emily"] = 7

		fmt.Println("familyAges value:", familyAges)

		// get the value of an existing key
		fmt.Println("John age:", familyAges["John"])

		// try to get the value of a none existing key
		fmt.Println("None existing key:", familyAges["None"])

		// change an existing key
		familyAges["Emily"] = 37
		fmt.Println("Emily age:", familyAges["Emily"])

		// delete a key
		delete(familyAges, "Jane")
		fmt.Println("familyAges value:", familyAges)

		// for each
		for name, age := range familyAges {
			fmt.Println("familyAges:", name, age)
		}

		// validate if a value exist
		_, ok := familyAges["Blunt"]
		if !ok {
			fmt.Println("familyAges['Blunt'] not found")
		}

		// delete all the keys on the map
		clear(familyAges)
		fmt.Println("familyAges value:", familyAges)
	}
}
