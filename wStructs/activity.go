package wStructs

import "fmt"

type Record struct {
	Name string
	Age  int
}

// string returns a csv representing our record.
func (r Record) String() string {
	return fmt.Sprintf("%s,%d", r.Name, r.Age)
}

func Activity(show bool) {
	if show {
		println("\n-- Structs Activity")

		david := Record{Name: "David Justice", Age: 28}
		sarah := Record{Name: "Sarah Murphy", Age: 28}
		fmt.Printf("%+v\n", david)
		fmt.Printf("%+v\n", sarah)

		john := Record{Name: "John Doak", Age: 100}
		fmt.Println(john.String())

		// changing a value
		david.Age = 77
		fmt.Println(david.String())
	}
}
