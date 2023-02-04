package main

import "fmt"

func main() {
	a := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(a)
	fmt.Println(a[:])
	fmt.Println(a[0])
	fmt.Println(a[0], a[1], a[2], a[3], a[4])
	fmt.Println(a[0:2])
	fmt.Println(a[1:4])
	fmt.Println(a[:2])
	fmt.Println(a[2:])

	slicingSlices()
	slicesShareMemory()
	appendOverlappingSlices()
}

func slicingSlices() {
	fmt.Println("--- slicing slices ---")
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:2]
	d := x[1:3]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func slicesShareMemory() {
	fmt.Println("--- slices share memory ---")
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]

	x[1] = 20
	y[0] = 10
	z[1] = 30

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func appendOverlappingSlices() {
	fmt.Println("--- append makes overlapping slices more confusing ---")
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Printf("x cap: %v, y cap: %v\n", cap(x), cap(y))
	fmt.Printf("x: %v, y: %v\n", x, y)

	y = append(y, 30)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
