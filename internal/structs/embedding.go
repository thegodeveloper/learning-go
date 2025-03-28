package structs

import "fmt"

type GDEmployee struct {
	Name string
	ID   string
}

func (e GDEmployee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	GDEmployee
	Reports []GDEmployee
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
