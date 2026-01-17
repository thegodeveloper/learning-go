package slices

import (
	"fmt"
)

func Run(show bool) {
	if show {
		fmt.Println("-- Slices")

		Definition(false)
		SlicingSlices(false)
		SlicesShareMemory(false)
		AppendOverlappingSlices(false)
		ConvertingArraysToSlices(false)
		CopySlice(false)
		CopySliceToArray(false)
		SliceWithSpecificValuesPerRun(false)
		SliceAppend(false)
		SliceWithMake(false)
		CompareSlices(false)
		EmptyingSlice(false)
		ChangeSlicing(false)
		ConfusingSlices(false)
		SlicesWithOverlappingStorage(false)
		PassMapSlice(false)
	}
}
