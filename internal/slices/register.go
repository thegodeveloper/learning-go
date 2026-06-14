package slices

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("slices", Run, map[string]func(bool){
		"definition":                    Definition,
		"slicingSlices":                 SlicingSlices,
		"slicesShareMemory":             SlicesShareMemory,
		"appendOverlappingSlices":       AppendOverlappingSlices,
		"convertingArraysToSlices":      ConvertingArraysToSlices,
		"copySlice":                     CopySlice,
		"copySliceToArray":              CopySliceToArray,
		"sliceWithSpecificValuesPerRun": SliceWithSpecificValuesPerRun,
		"sliceAppend":                   SliceAppend,
		"sliceWithMake":                 SliceWithMake,
		"compareSlices":                 CompareSlices,
		"emptyingSlice":                 EmptyingSlice,
		"changeSlicing":                 ChangeSlicing,
		"confusingSlices":               ConfusingSlices,
		"slicesWithOverlappingStorage":  SlicesWithOverlappingStorage,
		"passMapSlice":                  PassMapSlice,
		"multiDimensional":              MultiDimensional,
	}))
}
