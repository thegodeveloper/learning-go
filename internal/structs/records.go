package structs

import "fmt"

type Persona struct {
	Name string
	Age  int
}

func record(show bool) {
	if show {
		record := createPersona("Bill", 50)

		fmt.Println(record.String())

		fmt.Println("Name:", record.Name)
		fmt.Println("Age:", record.Age)
	}
}

func createPersona(name string, age int) Persona {
	return Persona{Name: name, Age: age}
}

// String returns a csv representating our record.
// This function receive a copy of the struct
func (r Persona) String() string {
	return fmt.Sprintf("%s,%d", r.Name, r.Age)
}
