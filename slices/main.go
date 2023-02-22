package main

import "fmt"

func main() {
	// this is a string array of 5 elements
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
	copyArrayToSlice()
	copySlice()
	copyingSliceToArray()
	sliceWithSpecificValuesPerIndex()
	sliceAppend()
}

func slicingSlices() {
	fmt.Println("--- slicing slices ---")
	// this is a slice
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

func sliceAppend() {
	var x []int
	fmt.Println("Initial Slice Length:", len(x))
	fmt.Println("Initial Slice Capacity:", cap(x))
	x = append(x, 10)
	fmt.Println("Slice Append, value of x:", x)
	fmt.Println("Slice Length:", len(x))
	fmt.Println("Slice Capacity:", cap(x))
	x = append(x, 20)
	fmt.Println("Slice Capacity:", cap(x))
	x = []int{10, 20, 30}
	fmt.Println("Current x values:", x)
	fmt.Println("New slice capacity:", cap(x))
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

func copyArrayToSlice() {
	fmt.Println("--- copy array to slice ---")
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[2:]
	x[0] = 10

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func copySlice() {
	fmt.Println("--- copying slice ---")
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)

	x = append(x, 7) // slices don't share memory
	x[0] = 77

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("number of elements copied:", num)

	fmt.Println("--- copying a slice is limitted by the smallest one ---")
	z := make([]int, 2)
	num = copy(z, x)
	fmt.Println("z:", z)
	fmt.Println("number of elements copied:", num)
}

func copyingSliceToArray() {
	fmt.Println("--- copying slice to array ---")
	x := []int{1, 2, 3, 4}  // slice
	d := [4]int{5, 6, 7, 8} // array
	y := make([]int, 2)
	copy(y, d[:]) // copying array to slice
	fmt.Println("y:", y)
	copy(d[:], x) // copying slice to array
	fmt.Println("d:", d)
}

func sliceWithSpecificValuesPerIndex() {
	// this creates a slice of 12 ints
	x := []int{1, 5: 6, 10: 100, 15}
	println("--- slice with specific values per index ---")
	println("x values: ", x)
}
