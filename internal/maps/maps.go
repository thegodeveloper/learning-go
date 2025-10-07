package maps

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Maps in Go")

		definition(false)

		mapsDeclarations(false)

		commaOkIdiomInMaps(false)

		deletingFromMaps(false)

		mapReadWrite(false)

		comparingMaps(false)

		mapSet(false)
	}
}
