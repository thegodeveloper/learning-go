package slices

import "fmt"

func Slicing(show bool) {
	if show {
		fmt.Println("--- Slicing ---")

		arr := [5]int{10, 20, 30, 40, 50}
		fmt.Println("Original array:", arr)

		// Create two slices from the same array.
		slice1 := arr[1:4]
		slice2 := arr[2:]
		fmt.Println("Slice1:", slice1)
		fmt.Println("Slice2:", slice2)

		fmt.Println("--- Modifying slice1 ---")
		slice1[1] = 777
		fmt.Println("Slice1 after modification:", slice1)
		fmt.Println("Slice2 after modification:", slice2)
		fmt.Println("Array after modification:", arr)
		fmt.Printf("\nSlice1 length: %d, capacity: %d\n", len(slice1), cap(slice1))

		fmt.Println("--- Appending within capacity ---")
		slice1 = append(slice1, 7777) // Still fits in the original array.
		fmt.Println("Slice1 after append:", slice1)
		fmt.Println("Array after append:", arr)

		fmt.Println("--- Appending beyond capacity ---")
		slice1 = append(slice1, 77777) // Now we exceed the capacity
		fmt.Println("Slice1 after exceeding capacity:", slice1)
		fmt.Println("Array after exceeding capacity:", arr)
		slice1[0] = 777777
		fmt.Println("Slice1 after independent modification:", slice1)
		fmt.Println("Array remains unchanged:", arr)
	}
}
