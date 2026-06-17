package structs

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// we are not modifying the receiver so that's why we are not using a pointer
// we are not handle nil
// we are using a value receiver
// by convention the first letter of the type, in this case letter p
// this is a method value receiver because we are not modifying the fields
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Go does not recommend to have Person and *Person on the same struct methods

// Capitalize Method to capitalize the Person string fields
// we are using here *Person because we are modifying the fields on the struct instance
// this is method with a pointer receiver because we are modifying the fields
func (p *Person) Capitalize() {
	p.FirstName = strings.ToUpper(p.FirstName)
	p.LastName = strings.ToUpper(p.LastName)
}

// IncrementAge increments the age field
// We use *Person because we are modifying the original struct instance
func (p *Person) IncrementAge() {
	p.Age++
}

func Methods(show bool) {
	if show {
		p := Person{
			FirstName: "Fred",
			LastName:  "Fredson",
			Age:       50,
		}

		output := p.String()
		fmt.Println("in Methods:", output)

		// methods on user defined type
		p.Capitalize()
		p.IncrementAge()
		fmt.Println("p.FirstName:", p.FirstName)
		fmt.Println("p.LastName:", p.LastName)
		fmt.Println("p.Age:", p.Age)
	}
}
