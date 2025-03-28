package interfaces

import "fmt"

// List the methods that must be implemented by a concrete type to meet the interface
type Stringer interface {
	String() string
}

type Persona struct {
	First, Last string
}

// A concrete type does not declare that it implements an interface.
// If the method set for a concrete type contains all the methods in the method set for the interface,
// the concrete type implements the interface.
// It is implemented implicity.
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
