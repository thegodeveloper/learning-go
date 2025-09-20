package structs

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// we are not modifying the receiver so that's why we are not using a pointer
// we are not handle nil
// we are using a value receiver
// by convention the first letter of the type, in this case letter p
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func methods(show bool) {
	if show {
		p := Person{
			FirstName: "Fred",
			LastName:  "Fredson",
			Age:       50,
		}

		output := p.String()
		fmt.Println("in Methods:", output)
	}
}
