package structs

import "fmt"

type Persona struct {
	Name string
	Age  int
}

func Records(show bool) {
	if show {
		record, err := NewPersona("Bill", 50)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(record.String())

		fmt.Println("Name:", record.Name)
		fmt.Println("Age:", record.Age)

		record.IncrAge()

		fmt.Println("New Age:", record.Age)
	}
}

// String returns a csv representating our record.
// This function receive a copy of the struct
func (r Persona) String() string {
	return fmt.Sprintf("%s,%d", r.Name, r.Age)
}

// IncrAge receive a pointer of the Persona struct and increase the Age field
func (r *Persona) IncrAge() {
	r.Age++
}

func NewPersona(name string, age int) (*Persona, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be the empty string")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age cannot be <= 0")
	}
	return &Persona{Name: name, Age: age}, nil
}
