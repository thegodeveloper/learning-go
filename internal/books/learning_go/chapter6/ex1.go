package chapter6

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func ex1(show bool) {
	if show {
		p1 := MakePerson("William", "Munoz", 50)
		fmt.Println(p1)

		p2 := MakePersonPointer("Magally", "Munoz", 47)
		fmt.Println(p2)
	}
}
