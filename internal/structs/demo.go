package structs

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
	pet  string
}

func demo(show bool) {
	if show {
		println("\n### Demo ###")

		structDeclaration()
		anonymousStruct()
	}
}

func structDeclaration() {
	fmt.Println("--- Struct Declaration ---")
	beth := person{
		age:  37,
		pet:  "dog",
		name: "Beth",
	}
	fmt.Println(beth)

	beth.name = "beth"
	fmt.Println("Name:", strings.ToTitle(beth.name))
	fmt.Println("Age:", beth.age)
	fmt.Println("Pet:", beth.pet)
}

func anonymousStruct() {
	fmt.Println("--- Anonymous Struct ---")
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println("Pet Value:", pet)
}
