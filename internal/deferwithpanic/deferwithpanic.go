package deferwithpanic

import "fmt"

func something() {
	defer fmt.Println("closed something")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i > 2 {
			panic("panic was called")
		}
	}
}

func Run(show bool) {
	if show {
		defer fmt.Println("close main")
		something()
		fmt.Println("something was finished") // this is not called
	}
}
