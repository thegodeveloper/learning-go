package main

import (
	"github.com/thegodeveloper/learning-go/internal/arrays"
	"github.com/thegodeveloper/learning-go/internal/books/learning_go/chapter5"
	"github.com/thegodeveloper/learning-go/internal/books/learning_go/chapter6"
	"github.com/thegodeveloper/learning-go/internal/box_ui"
	"github.com/thegodeveloper/learning-go/internal/channels"
	"github.com/thegodeveloper/learning-go/internal/closures"
	"github.com/thegodeveloper/learning-go/internal/concurrency"
	"github.com/thegodeveloper/learning-go/internal/constants"
	"github.com/thegodeveloper/learning-go/internal/deferpanic"
	"github.com/thegodeveloper/learning-go/internal/deferpanicrecover"
	"github.com/thegodeveloper/learning-go/internal/defersimplecat"
	"github.com/thegodeveloper/learning-go/internal/deferwithpanic"
	"github.com/thegodeveloper/learning-go/internal/enums"
	"github.com/thegodeveloper/learning-go/internal/forloops"
	"github.com/thegodeveloper/learning-go/internal/functionaloptions"
	"github.com/thegodeveloper/learning-go/internal/functions"
	"github.com/thegodeveloper/learning-go/internal/goroutines"
	"github.com/thegodeveloper/learning-go/internal/gotousecase"
	"github.com/thegodeveloper/learning-go/internal/helloworld"
	"github.com/thegodeveloper/learning-go/internal/input"
	"github.com/thegodeveloper/learning-go/internal/interfaces"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/escapeanalysis"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/memoryprofiling"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/stacksandpointers/frameboundaries"
	"github.com/thegodeveloper/learning-go/internal/languagemechanics/stacksandpointers/indirectmemoryaccess"
	"github.com/thegodeveloper/learning-go/internal/lengthcapacity"
	"github.com/thegodeveloper/learning-go/internal/logfatal"
	"github.com/thegodeveloper/learning-go/internal/loopsbranches"
	"github.com/thegodeveloper/learning-go/internal/maps"
	"github.com/thegodeveloper/learning-go/internal/openfile"
	"github.com/thegodeveloper/learning-go/internal/passingarguments"
	"github.com/thegodeveloper/learning-go/internal/pointers"
	"github.com/thegodeveloper/learning-go/internal/printstars"
	"github.com/thegodeveloper/learning-go/internal/random_message"
	"github.com/thegodeveloper/learning-go/internal/runes"
	"github.com/thegodeveloper/learning-go/internal/sha256"
	"github.com/thegodeveloper/learning-go/internal/slices"
	"github.com/thegodeveloper/learning-go/internal/statistics"
	"github.com/thegodeveloper/learning-go/internal/strings"
	"github.com/thegodeveloper/learning-go/internal/structs"
	"github.com/thegodeveloper/learning-go/internal/sumtwonumbers"
	"github.com/thegodeveloper/learning-go/internal/switches"
	"github.com/thegodeveloper/learning-go/internal/syslog"
	"github.com/thegodeveloper/learning-go/internal/typediffarrayslicesitem"
	"github.com/thegodeveloper/learning-go/internal/unicode"
	"github.com/thegodeveloper/learning-go/internal/values"
	"github.com/thegodeveloper/learning-go/internal/variables"
	"github.com/thegodeveloper/learning-go/internal/wadapter"
	"github.com/thegodeveloper/learning-go/internal/wawspresignedurl"
	"github.com/thegodeveloper/learning-go/internal/wcalculator"
	"github.com/thegodeveloper/learning-go/internal/wcomplex"
	"github.com/thegodeveloper/learning-go/internal/wconstants"
	"github.com/thegodeveloper/learning-go/internal/wcountoccurrences"
	"github.com/thegodeveloper/learning-go/internal/wcustomtype"
	"github.com/thegodeveloper/learning-go/internal/wdefer"
	"github.com/thegodeveloper/learning-go/internal/webrouting"
	"github.com/thegodeveloper/learning-go/internal/werrorold"
	"github.com/thegodeveloper/learning-go/internal/werrors"
	"github.com/thegodeveloper/learning-go/internal/wexternalpackages"
	"github.com/thegodeveloper/learning-go/internal/wgoroutines"
	"github.com/thegodeveloper/learning-go/internal/which"
	"github.com/thegodeveloper/learning-go/internal/wintegertypes"
	"github.com/thegodeveloper/learning-go/internal/wjson"
	"github.com/thegodeveloper/learning-go/internal/wnilzero"
	"github.com/thegodeveloper/learning-go/internal/wopenpage"
	"github.com/thegodeveloper/learning-go/internal/wpackages"
	"github.com/thegodeveloper/learning-go/internal/wrand"
	"github.com/thegodeveloper/learning-go/internal/wtime"
	"github.com/thegodeveloper/learning-go/internal/wtypeconversions"
	"github.com/thegodeveloper/logging"
)

/*func init() {
	fmt.Println("init from my program")
}*/

func main() {
	box_ui.Master(false)
	switches.Master(false)
	unicode.Master(false)
	values.Master(false)
	strings.Master(false)
	slices.Master(false)
	sha256.Master(false)
	runes.Master(false)
	pointers.Master(false)
	maps.Master(false)
	variables.Master(false)
	typediffarrayslicesitem.Master(false)
	sumtwonumbers.Master(false)
	passingarguments.Master(false)
	openfile.Master(false)
	loopsbranches.Master(false)
	lengthcapacity.Master(false)
	indirectmemoryaccess.Master(false)
	frameboundaries.Master(false)
	escapeanalysis.Master(false)
	memoryprofiling.Master(false)
	helloworld.Master(false)
	functions.Master(false)
	functionaloptions.Master(false)
	deferwithpanic.Master(false)
	forloops.Master(false)
	enums.Master(false)
	deferpanicrecover.Master(false)
	werrors.Master(false)
	deferpanic.Master(false)
	arrays.Master(false)
	interfaces.Master(false)
	wrand.Master(false)
	wpackages.Master(false)
	wexternalpackages.Master(false)
	wadapter.Master(false)
	channels.Master(false)
	goroutines.Master(false)
	concurrency.Master(false)
	structs.Master(false)
	wopenpage.Master(false)
	wawspresignedurl.Master(false)
	wjson.Master(false)
	wtime.Master(false)
	wcalculator.Master(false)
	wcomplex.Master(false)
	wconstants.Master(false)
	wcountoccurrences.Master(false)
	wcustomtype.Master(false)
	wdefer.Master(false)
	wnilzero.Master(false)
	werrorold.Master(false)
	wgoroutines.Master(false)
	wintegertypes.Master(false)
	wtypeconversions.Master(false)
	printstars.Master(false)
	gotousecase.Master(false)
	closures.Master(false)
	defersimplecat.Master(false)
	constants.Master(false)
	input.Master(false)
	which.Master(false)
	syslog.Master(false)
	logfatal.Master(false)
	logging.Master(false, true, false)
	statistics.Master(false)
	randommessage.Index(false)

	// Learning Go Exercises
	chapter5.Index(false)
	chapter6.Index(false)

	webrouting.Index(false)
}
