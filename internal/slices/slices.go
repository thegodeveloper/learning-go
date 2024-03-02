package slices

import (
	"fmt"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Slices")

		sliceDefinition(false)
		slicingSlices(false)
		slicesShareMemory(false)
		appendOverlappingSlices(false)
		copyArrayToSlice(false)
		copySlice(false)
		copySliceToArray(false)
		sliceWithSpecificValuesPerIndex(false)
		sliceAppend(false)
		sliceWithMake(false)

		err := slicesAsBuffers(false, "./books.txt") // this is wrong
		if err != nil {
			fmt.Println(err)
		}

		compareSlices(false)
		emptyingSlice(false)
		changeSlicing(false)
	}
}
