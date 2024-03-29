package escapeanalysis

import "fmt"

type user struct {
	name  string
	email string
}

func Master(show bool) {
	if show {
		u1 := createUserV1()
		u2 := createUserV2()

		fmt.Println("u1", &u1, "u2", &u2)
	}
}

//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	fmt.Println("V1", &u)
	return u
}

//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	fmt.Println("V2", &u)
	return &u
}
