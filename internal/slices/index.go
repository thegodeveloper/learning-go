package slices

import (
	"fmt"
)

func Master(show bool) {
	if show {
		fmt.Println("-- Slices")

		definition(false)
		slicingSlices(false)
		slicesShareMemory(false)
		appendOverlappingSlices(false)
		convertingArraysToSlices(false)
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
		confusingSlices(false)
		slicesWithOverlappingStorage(false)
		passMapSlice(false)
	}
}
