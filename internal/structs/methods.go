package structs

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
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
	}
}
