package functions

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func sumVariadic(nums ...int) int {
	total := 0
	for _, a := range nums {
		total = total + a
	}
	return total
}

func ops(a int, b int) (int, int) {
	return a + b, a - b
}

func doit(operator func(int, int) int, a int, b int) int {
	return operator(a, b)
}

func accumulator(increment int) func() int {
	i := 0
	return func() int {
		i = i + increment
		return i
	}
}

func Master(show bool) {
	if show {
		fmt.Println("--- simulating named and optional parameters ---")
		simulatingNamedOptionalParams(MyFuncOpts{
			LastName: "Patel",
			Age:      50,
		})
		simulatingNamedOptionalParams(MyFuncOpts{
			FirstName: "Joe",
			LastName:  "Smith",
		})

		fmt.Println("--- using a variadic function ---")
		fmt.Println(addTo(3))
		fmt.Println(addTo(3, 2))
		fmt.Println(addTo(3, 5, 7, 11, 13, 17))
		a := []int{4, 3}
		fmt.Println(addTo(7, a...))
		fmt.Println(addTo(7, []int{1, 3, 5, 7, 11, 13, 17}...))

		// Returning multiple values
		result, remainder, err := divAndRemainder(7, 7)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("result: %d, remainder: %d", result, remainder)

		result, remainder, err = divAndRemainderNamed(7, 7)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("result: %d, remainder: %d\n", result, remainder)

		sortSlice()

		fmt.Println("--- return a function ---")
		twoBase := makeMult(2)
		threeBase := makeMult(3)
		for i := 0; i < 3; i++ {
			fmt.Println(twoBase(i), threeBase(i))
		}

		fmt.Println("--- print file content ---")
		printFileContent("./books.txt")
	}
}
