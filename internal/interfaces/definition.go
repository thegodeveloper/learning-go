package interfaces

import "fmt"

type Stringer interface {
	String() string
}

type Persona struct {
	First, Last string
}

func (p Persona) String() string {
	return fmt.Sprintf("%s %s", p.First, p.Last)
}

func definition() {
	p := Persona{
		First: "Will",
		Last:  "MR",
	}

	fmt.Println("Full Name: ", p.String())
}
