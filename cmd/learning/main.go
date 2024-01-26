package main

import (
	"github.com/thegodeveloper/learning-go/internal/channels"
	"github.com/thegodeveloper/learning-go/internal/deferpanic"
	"github.com/thegodeveloper/learning-go/internal/deferpanicrecover"
	"github.com/thegodeveloper/learning-go/internal/deferwithpanic"
	"github.com/thegodeveloper/learning-go/internal/enums"
	"github.com/thegodeveloper/learning-go/internal/forloops"
	"github.com/thegodeveloper/learning-go/internal/functionaloptions"
	"github.com/thegodeveloper/learning-go/internal/functions"
	"github.com/thegodeveloper/learning-go/internal/goroutines"
	"github.com/thegodeveloper/learning-go/internal/helloworld"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/escapeanalysis"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/memoryprofiling"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/stacksandpointers/frameboundaries"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/stacksandpointers/indirectmemoryaccess"
	"github.com/thegodeveloper/learning-go/internal/lengthcapacity"
	"github.com/thegodeveloper/learning-go/internal/loopsbranches"
	"github.com/thegodeveloper/learning-go/internal/maps"
	"github.com/thegodeveloper/learning-go/internal/openfile"
	"github.com/thegodeveloper/learning-go/internal/passingarguments"
	"github.com/thegodeveloper/learning-go/internal/pointers"
	"github.com/thegodeveloper/learning-go/internal/runes"
	"github.com/thegodeveloper/learning-go/internal/sha256"
	"github.com/thegodeveloper/learning-go/internal/slices"
	"github.com/thegodeveloper/learning-go/internal/sumtwonumbers"
	"github.com/thegodeveloper/learning-go/internal/typediffarrayslicesitem"
	"github.com/thegodeveloper/learning-go/internal/variables"
	"github.com/thegodeveloper/learning-go/wAWSPresignedURL"
	"github.com/thegodeveloper/learning-go/wAdapter"
	"github.com/thegodeveloper/learning-go/wArrays"
	"github.com/thegodeveloper/learning-go/wCalculator"
	"github.com/thegodeveloper/learning-go/wComplex"
	"github.com/thegodeveloper/learning-go/wConcurrency"
	"github.com/thegodeveloper/learning-go/wConstants"
	"github.com/thegodeveloper/learning-go/wCountOccurrencesInFile"
	"github.com/thegodeveloper/learning-go/wCustomType"
	"github.com/thegodeveloper/learning-go/wDefer"
	"github.com/thegodeveloper/learning-go/wErrorOld"
	"github.com/thegodeveloper/learning-go/wErrors"
	"github.com/thegodeveloper/learning-go/wExternalPackages"
	"github.com/thegodeveloper/learning-go/wGoroutines"
	"github.com/thegodeveloper/learning-go/wIntegerTypes"
	wIntefaces "github.com/thegodeveloper/learning-go/wInterfaces"
	"github.com/thegodeveloper/learning-go/wJSON"
	"github.com/thegodeveloper/learning-go/wNilZero"
	"github.com/thegodeveloper/learning-go/wOpenPage"
	"github.com/thegodeveloper/learning-go/wPackages"
	"github.com/thegodeveloper/learning-go/wRand"
	"github.com/thegodeveloper/learning-go/wStructs"
	"github.com/thegodeveloper/learning-go/wTest"
	"github.com/thegodeveloper/learning-go/wTime"
)

/*func init() {
	fmt.Println("init from my program")
}*/

func main() {
	// Slices
	slices.Master(false)

	// sha256
	sha256.Master(false)

	// Runes
	runes.Master(false)

	// Pointers
	pointers.Master(false)

	// Maps
	maps.Master(false)

	// Variables
	variables.Master(false)

	// Type Diff Array Slices Item
	typediffarrayslicesitem.Master(false)

	// Sum two Arguments
	sumtwonumbers.Master(false)

	// Passing Arguments
	passingarguments.Master(false)

	// Open File
	openfile.Master(false)

	// Loops Branches
	loopsbranches.Master(false)

	// Length Capacity
	lengthcapacity.Master(false)

	// Indirect Memory Access
	indirectmemoryaccess.Master(false)

	// Frame Boundaries
	frameboundaries.Master(false)

	// Escape Analysis
	escapeanalysis.Master(false)

	// Memory Profiling
	memoryprofiling.Master(false)

	// Hello World
	helloworld.Master(false)

	// Functions
	functions.Master(false)

	// Functional Options Pattern
	functionaloptions.Master(false)

	// Defer with Panic
	deferwithpanic.Master(false)

	// For Loops
	forloops.Master(false)

	// Enums
	enums.Master(false)

	// Defer Panic Recover
	deferpanicrecover.Master(false)

	// Errors
	wErrors.Master(false)

	// Defer Panic
	deferpanic.Master(false)

	// Arrays
	wArrays.Master(false)

	// Interfaces
	wIntefaces.Master(false)

	// Random Values
	wRand.Master(false)

	// Packages
	wPackages.Master(false)

	// External Packages
	wExternalPackages.Master(false)

	// Adapter Pattern
	wAdapter.Master(false)

	// Channels
	channels.Master(false)

	// Goroutines
	goroutines.Master(false)

	// Concurrency
	wConcurrency.Master(false)

	// Structs
	wStructs.Master(false)

	// Open Web Page
	wOpenPage.Master(false)

	// AWS Pre-signed URL
	wAWSPresignedURL.Master(false)

	// JSON
	wJSON.Master(false)

	// Time
	wTime.Master(false)

	// Test
	wTest.Master(false)

	// Calculator
	wCalculator.Master(false)

	// Complex
	wComplex.Master(false)

	// Constants
	wConstants.Master(false)

	// Occurrences in File
	wCountOccurrencesInFile.Master(false)

	// Custom Type
	wCustomType.Master(false)

	// Defer
	wDefer.Master(false)

	// Nil and Zero Values
	wNilZero.Master(false)

	// Errors Old
	wErrorOld.Master(false)

	// Goroutines Example
	wGoroutines.Master(false)

	// Integer Types
	wIntegerTypes.Master(false)
}
