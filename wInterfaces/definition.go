package wIntefaces

import "fmt"

type Stringer interface {
	String() string
}

type Person struct {
	First, Last string
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s", p.First, p.Last)
}

func definition() {
	p := Person{
		First: "Will",
		Last:  "MR",
	}

	fmt.Println("Full Name: ", p.String())
}
