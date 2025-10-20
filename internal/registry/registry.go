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

func Names() []string {
	names := make([]string, 0, len(Packages))
	for name := range Packages {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
