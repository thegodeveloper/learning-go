package structs

import "fmt"

type GDEmployee struct {
	Name string
	ID   string
}

func (e GDEmployee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

// You can embed any type within a struct, not just another struct.
// This promotes the methods on the embedded type to the containing struct.
// Embedding is not inheritance.
type Manager struct {
	GDEmployee
	Reports []GDEmployee
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// Here we are embedding two interfaces in one
type ReadCloser interface {
	Reader
	Closer
}

func embedding(show bool) {
	if show {
		m := Manager{
			GDEmployee: GDEmployee{
				Name: "Bob Bobson",
				ID:   "12345",
			},
			Reports: []GDEmployee{},
		}
		fmt.Println(m.ID)
		fmt.Println(m.Description())
	}
}
