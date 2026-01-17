package registry

import (
	"sort"

	"github.com/thegodeveloper/learning-go/internal/adapter"
	"github.com/thegodeveloper/learning-go/internal/arrays"
	"github.com/thegodeveloper/learning-go/internal/awspresignedurl"
	"github.com/thegodeveloper/learning-go/internal/basics"
	"github.com/thegodeveloper/learning-go/internal/calculator"
	"github.com/thegodeveloper/learning-go/internal/channels"
	"github.com/thegodeveloper/learning-go/internal/closures"
	"github.com/thegodeveloper/learning-go/internal/complexnumbers"
	"github.com/thegodeveloper/learning-go/internal/concurrency"
	"github.com/thegodeveloper/learning-go/internal/constants"
	"github.com/thegodeveloper/learning-go/internal/countoccurrences"
	"github.com/thegodeveloper/learning-go/internal/customtype"
	"github.com/thegodeveloper/learning-go/internal/deferpanic"
	"github.com/thegodeveloper/learning-go/internal/defersimplecat"
	"github.com/thegodeveloper/learning-go/internal/deferstatement"
	"github.com/thegodeveloper/learning-go/internal/deferwithpanic"
	"github.com/thegodeveloper/learning-go/internal/enums"
	"github.com/thegodeveloper/learning-go/internal/errorold"
	"github.com/thegodeveloper/learning-go/internal/errors"
	"github.com/thegodeveloper/learning-go/internal/externalpackages"
	"github.com/thegodeveloper/learning-go/internal/forloops"
	"github.com/thegodeveloper/learning-go/internal/functionaloptions"
	"github.com/thegodeveloper/learning-go/internal/functions"
	"github.com/thegodeveloper/learning-go/internal/generics"
	"github.com/thegodeveloper/learning-go/internal/goroutines"
	"github.com/thegodeveloper/learning-go/internal/gotousecase"
	"github.com/thegodeveloper/learning-go/internal/helloworld"
	"github.com/thegodeveloper/learning-go/internal/input"
	"github.com/thegodeveloper/learning-go/internal/integertypes"
	"github.com/thegodeveloper/learning-go/internal/interfaces"
	"github.com/thegodeveloper/learning-go/internal/json"
	"github.com/thegodeveloper/learning-go/internal/lengthcapacity"
	"github.com/thegodeveloper/learning-go/internal/logfatal"
	"github.com/thegodeveloper/learning-go/internal/loopsbranches"
	"github.com/thegodeveloper/learning-go/internal/maps"
	"github.com/thegodeveloper/learning-go/internal/nilzero"
	"github.com/thegodeveloper/learning-go/internal/openfile"
	"github.com/thegodeveloper/learning-go/internal/openpage"
	"github.com/thegodeveloper/learning-go/internal/packages"
	"github.com/thegodeveloper/learning-go/internal/passingarguments"
	"github.com/thegodeveloper/learning-go/internal/pointers"
	"github.com/thegodeveloper/learning-go/internal/rand"
	"github.com/thegodeveloper/learning-go/internal/randommessage"
	"github.com/thegodeveloper/learning-go/internal/reflection"
	"github.com/thegodeveloper/learning-go/internal/runes"
	"github.com/thegodeveloper/learning-go/internal/sha256"
	"github.com/thegodeveloper/learning-go/internal/slices"
	"github.com/thegodeveloper/learning-go/internal/statistics"
	"github.com/thegodeveloper/learning-go/internal/strings"
	"github.com/thegodeveloper/learning-go/internal/structs"
	"github.com/thegodeveloper/learning-go/internal/sumtwonumbers"
	"github.com/thegodeveloper/learning-go/internal/switches"
	"github.com/thegodeveloper/learning-go/internal/syslog"
	"github.com/thegodeveloper/learning-go/internal/time"
	"github.com/thegodeveloper/learning-go/internal/typeconversions"
	"github.com/thegodeveloper/learning-go/internal/typediffarrayslicesitem"
	"github.com/thegodeveloper/learning-go/internal/unicode"
	"github.com/thegodeveloper/learning-go/internal/values"
	"github.com/thegodeveloper/learning-go/internal/variables"
	"github.com/thegodeveloper/learning-go/internal/webrouting"
	"github.com/thegodeveloper/learning-go/internal/which"
)

// PackageFunctions maps function names to their execution functions within a package
type PackageFunctions map[string]func(bool)

var Packages = map[string]func(bool){
	"adapter":                 adapter.Run,
	"arrays":                  arrays.Run,
	"awspresignedurl":         awspresignedurl.Run,
	"basics":                  basics.Run,
	"calculator":              calculator.Run,
	"channels":                channels.Run,
	"closures":                closures.Run,
	"complexnumbers":          complexnumbers.Run,
	"concurrency":             concurrency.Run,
	"constants":               constants.Run,
	"countoccurrences":        countoccurrences.Run,
	"customtype":              customtype.Run,
	"deferpanic":              deferpanic.Run,
	"deferpanicrecover":       deferpanic.Run,
	"defersimplecat":          defersimplecat.Run,
	"deferstatement":          deferstatement.Run,
	"deferwithpanic":          deferwithpanic.Run,
	"enums":                   enums.Run,
	"errorold":                errorold.Run,
	"errors":                  errors.Run,
	"externalpackages":        externalpackages.Run,
	"forloops":                forloops.Run,
	"functionaloptions":       functionaloptions.Run,
	"functions":               functions.Run,
	"generics":                generics.Run,
	"goroutines":              goroutines.Run,
	"gotousecase":             gotousecase.Run,
	"helloworld":              helloworld.Run,
	"input":                   input.Run,
	"integertypes":            integertypes.Run,
	"interfaces":              interfaces.Run,
	"json":                    json.Run,
	"lengthcapacity":          lengthcapacity.Run,
	"logfatal":                logfatal.Run,
	"loopsbranches":           loopsbranches.Run,
	"maps":                    maps.Run,
	"nilzero":                 nilzero.Run,
	"openfile":                openfile.Run,
	"openpage":                openpage.Run,
	"packages":                packages.Run,
	"passingarguments":        passingarguments.Run,
	"pointers":                pointers.Run,
	"rand":                    rand.Run,
	"randommessage":           randommessage.Run,
	"reflection":              reflection.Run,
	"runes":                   runes.Run,
	"sha256":                  sha256.Run,
	"slices":                  slices.Run,
	"statistics":              statistics.Run,
	"strings":                 strings.Run,
	"structs":                 structs.Run,
	"sumtwonumbers":           sumtwonumbers.Run,
	"switches":                switches.Run,
	"syslog":                  syslog.Run,
	"time":                    time.Run,
	"typeconversions":         typeconversions.Run,
	"typediffarrayslicesitem": typediffarrayslicesitem.Run,
	"unicode":                 unicode.Run,
	"values":                  values.Run,
	"variables":               variables.Run,
	"webrouting":              webrouting.Run,
	"which":                   which.Run,
}

// PackageRegistry maps package names to their sub-functions
// Packages that have entries here will show their functions instead of running directly
var PackageRegistry = map[string]PackageFunctions{
	"adapter": {
		"macExample":     adapter.MacExample,
		"windowsExample": adapter.WindowsExample,
	},
	"arrays": {
		"definition":              arrays.ArrayDefinition,
		"basics":                  arrays.Basics,
		"symbols":                 arrays.Symbols,
		"printIndexesAndElements": arrays.PrintIndexesAndElements,
		"printType":               arrays.PrintType,
		"comparingArrays":         arrays.ComparingArrays,
		"arrayFunctions":          arrays.ArrayFunctions,
	},
	"basics": {
		"net": basics.Net,
	},
	"channels": {
		"bufferedChannel":              channels.BufferedChannel,
		"unbufferedChannel":            channels.UnbufferedChannel,
		"readAndWriteToFromChannel":    channels.ReadAndWriteToFromChannel,
		"closedChannel":                channels.ClosedChannel,
		"channelsAsFunctionParameters": channels.ChannelsAsFunctionParameters,
		"definition":                   channels.Definition,
	},
	"closures": {
		"declaration": closures.Declaration,
	},
	"errorold": {
		"demo":              errorold.Demo,
		"errorsAndValues":   errorold.ErrorsAndValues,
		"errorFromFunction": errorold.ErrorFromFunction,
		"wrappingErrors":    errorold.WrappingErrors,
		"simple":            errorold.Simple,
		"sentinelError":     errorold.SentinelError,
	},
	"errors": {
		"basics":          errors.Basics,
		"simpleErrors":    errors.SimpleErrors,
		"sentinelErrors":  errors.SentinelErrors,
		"errorsAreValues": errors.ErrorsAreValues,
		"wrappingErrors":  errors.WrappingErrors,
		"errorsIs":        errors.ErrorsIs,
		"namedErrors":     errors.NamedErrors,
		"allExamples":     errors.AllExamples,
	},
	"forloops": {
		"definition":         forloops.Definition,
		"forContinue":        forloops.ForContinue,
		"forRange":           forloops.ForRange,
		"forRangeMaps":       forloops.ForRangeMaps,
		"forString":          forloops.ForString,
		"forWithLabels":      forloops.ForWithLabels,
		"forBreakFromSwitch": forloops.ForBreakFromSwitch,
		"rangeIsACopy":       forloops.RangeIsACopy,
	},
	"functions": {
		"simulatingOptionalParams": functions.SimulatingOptionalParams,
		"variadicParameters":       functions.VariadicParameters,
		"multipleReturnValues":     functions.MultipleReturnValues,
		"namedReturnValues":        functions.NamedReturnValues,
		"sortSlice":                functions.SortSlice,
		"returnFunction":           functions.ReturnFunction,
		"printFileContent":         functions.PrintFileContent,
		"functionsAreValues":       functions.FunctionsAreValues,
		"anonymousFunctions":       functions.AnonymousFunctions,
		"functionsAsParameters":    functions.FunctionsAsParameters,
		"prefixer":                 functions.Prefixer,
	},
	"generics": {
		"introduction": generics.Introduction,
	},
	"goroutines": {
		"implementCountTo":                          goroutines.ImplementCountTo,
		"implementCancelFunctionTerminateGoroutine": goroutines.ImplementCancelFunctionTerminateGoroutine,
		"implementBackPressure":                     goroutines.ImplementBackPressure,
		"passingCopyGoroutine":                      goroutines.PassingCopyGoroutine,
		"allExamples":                               goroutines.AllExamples,
	},
	"interfaces": {
		"definition":         interfaces.Definition,
		"zeroValue":          interfaces.ZeroValue,
		"emptyInterface":     interfaces.EmptyInterface,
		"implementInterface": interfaces.ImplementInterface,
		"implicit":           interfaces.Implicit,
	},
	"maps": {
		"definition":   maps.Definition,
		"declarations": maps.MapsDeclarations,
		"commaOkIdiom": maps.CommaOkIdiomInMaps,
		"deleting":     maps.DeletingFromMaps,
		"readWrite":    maps.MapReadWrite,
		"comparing":    maps.ComparingMaps,
		"set":          maps.MapSet,
	},
	"pointers": {
		"declaration": pointers.Declaration,
	},
	"slices": {
		"definition":                    slices.Definition,
		"slicingSlices":                 slices.SlicingSlices,
		"slicesShareMemory":             slices.SlicesShareMemory,
		"appendOverlappingSlices":       slices.AppendOverlappingSlices,
		"convertingArraysToSlices":      slices.ConvertingArraysToSlices,
		"copySlice":                     slices.CopySlice,
		"copySliceToArray":              slices.CopySliceToArray,
		"sliceWithSpecificValuesPerRun": slices.SliceWithSpecificValuesPerRun,
		"sliceAppend":                   slices.SliceAppend,
		"sliceWithMake":                 slices.SliceWithMake,
		"compareSlices":                 slices.CompareSlices,
		"emptyingSlice":                 slices.EmptyingSlice,
		"changeSlicing":                 slices.ChangeSlicing,
		"confusingSlices":               slices.ConfusingSlices,
		"slicesWithOverlappingStorage":  slices.SlicesWithOverlappingStorage,
		"passMapSlice":                  slices.PassMapSlice,
	},
	"strings": {
		"extractingSingleValue": strings.ExtractingSingleValue,
		"toBeAwareInStrings":    strings.ToBeAwareInStrings,
	},
	"structs": {
		"definition":      structs.Definition,
		"demo":            structs.Demo,
		"activity":        structs.Activity,
		"records":         structs.Records,
		"pointerReceiver": structs.PointerReceiver,
		"methods":         structs.Methods,
		"binaryTree":      structs.BinaryTree,
		"songs":           structs.Songs,
		"embedding":       structs.Embedding,
	},
}

func Names() []string {
	names := make([]string, 0, len(Packages))
	for name := range Packages {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// GetPackageFunctions returns the functions for a package, or nil if it doesn't have sub-functions
func GetPackageFunctions(packageName string) PackageFunctions {
	return PackageRegistry[packageName]
}

// GetFunctionNames returns sorted function names for a package
func GetFunctionNames(packageName string) []string {
	funcs := PackageRegistry[packageName]
	if funcs == nil {
		return []string{}
	}

	names := make([]string, 0, len(funcs))
	for name := range funcs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// HasSubFunctions checks if a package has registered sub-functions
func HasSubFunctions(packageName string) bool {
	funcs, exists := PackageRegistry[packageName]
	return exists && len(funcs) > 0
}
