package wStructs

import (
	"fmt"
	"time"
)

func Master(show bool) {
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

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       string
	Position  string
	Salary    int
	ManagerID int
}
