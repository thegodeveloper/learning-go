package structs

import "fmt"

type Persona struct {
	Name string
	Age  int
}

func record(show bool) {
	if show {
		record := createPersona("Bill", 50)

		fmt.Println("Name:", record.Name)
		fmt.Println("Age:", record.Age)
	}
}

func createPersona(name string, age int) Persona {
	return Persona{Name: name, Age: age}
}
