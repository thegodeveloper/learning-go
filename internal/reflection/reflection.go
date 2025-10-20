package reflection

import (
	"fmt"
	"reflect"
)

type UserService struct {
	Name string
}

func (UserService) CreateUser(email string) string {
	return fmt.Sprintf("User with email %s created", email)
}

func (UserService) DeleteUser(id int) string {
	return fmt.Sprintf("User with ID %d deleted", id)
}

func Run(show bool) {
	if show {
		fmt.Println("Reflection in Go")

		svc := UserService{Name: "AccountService"}

		t := reflect.TypeOf(svc)
		v := reflect.ValueOf(svc)

		fmt.Println("🔍 Struct name:", t.Name())
		fmt.Println("🔍 Number of methods:", t.NumMethod())
		fmt.Println("")

		// List methods dynamically
		for i := 0; i < t.NumMethod(); i++ {
			m := t.Method(i)
			fmt.Printf("➡️  Method #%d: %s (Type: %s)\n", i, m.Name, m.Type)
		}
		fmt.Println("")

		// -------------------------------------------------------------------------
		// Dynamically call a method by name using reflection
		// -------------------------------------------------------------------------

		method := v.MethodByName("CreateUser")
		if method.IsValid() {
			args := []reflect.Value{reflect.ValueOf("foo@bar.com")}
			results := method.Call(args)

			fmt.Println("🚀 Dynamic call result:", results[0].Interface())
		}

		method = v.MethodByName("DeleteUser")
		if method.IsValid() {
			args := []reflect.Value{reflect.ValueOf(7)}
			results := method.Call(args)

			fmt.Println("🚀 Dynamic call result:", results[0].Interface())
		}
	}
}
