package deferpanicrecover

import "fmt"

func something() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("No need to panic if i = ", r)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i > 2 {
			panic(i)
		}
	}
	fmt.Println("closed something normally")
}

func Index(show bool) {
	if show {
		defer fmt.Println("closed main")
		something()
		fmt.Println("main was finished")
	}
}
