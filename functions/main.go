package main

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

func main() {
	res := sum(7, 7)
	fmt.Println(res)

	sum, subs := ops(7, 7)
	fmt.Println("7 + 7 = ", sum, "7 - 7 =", subs)
	b, _ := ops(7, 7)
	fmt.Println("7 + 7=", b)

	// calling sum variadic number
	total := sumVariadic(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("the first five numbers sum is", total)

	c := doit(sum, 7, 7)
	fmt.Println("7 + 7=", c)
	d := doit(multiply, 7, 7)
	fmt.Println("2 * 3 =", d)

	a := accumulator(1)
	b := accumulator(2)

	fmt.Println("a", "b")
	for i:= 0; i < 5; i++ {
		fmt.Println(a(), b(()))
	}
}
