package registry

import (
	"sort"

	"github.com/thegodeveloper/learning-go/internal/adapter"
	"github.com/thegodeveloper/learning-go/internal/arrays"
	"github.com/thegodeveloper/learning-go/internal/awspresignedurl"
	"github.com/thegodeveloper/learning-go/internal/basics"
	"github.com/thegodeveloper/learning-go/internal/concurrency"
)

var Packages = map[string]func(bool){
	"adapter":         adapter.Run,
	"arrays":          arrays.Run,
	"awspresignedurl": awspresignedurl.Run,
	"basics":          basics.Run,
	"concurrency":     concurrency.Run,
}

func Names() []string {
	names := make([]string, 0, len(Packages))
	for name := range Packages {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
