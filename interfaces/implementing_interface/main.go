package main

import "fmt"

type Person interface {
	greet() string
}

type Human struct {
	Name string
}

func (h *Human) greet() string {
	return "Hi, I'm " + h.Name
}

func isAPerson(h Person) {
	fmt.Println(h.greet())
}

func main() {
	var a = Human{
		Name: "Will",
	}

	fmt.Println(a.greet())

	isAPerson(&a)
}
