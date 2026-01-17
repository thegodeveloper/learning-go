package structs

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       string
	Position  string
	Salary    int
	ManagerID int
}

type tree struct {
	value       int
	left, right *tree
}

func Definition(show bool) {
	if show {
		fmt.Println("-- Structs")
		bill := createEmployee(
			123,
			"Bill",
			"Street 7",
			time.DateOnly,
			"Professional Go Developer",
			13000,
			123)
		fmt.Println(bill)

		a := []int{777, 7, 77, 77777, 7777}
		fmt.Printf("Initial value of a: %v\n", a)
		Sort(a)
		fmt.Printf("Final value of a: %v\n", a)
	}
}

func createEmployee(id int, name string, address string, dob string, position string, salary int, managerId int) Employee {
	return Employee{
		ID:        id,
		Name:      name,
		Address:   address,
		DoB:       dob,
		Position:  position,
		Salary:    salary,
		ManagerID: managerId,
	}
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
