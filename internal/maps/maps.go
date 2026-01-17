package maps

import "fmt"

func Run(show bool) {
	if show {
		fmt.Println("-- Maps in Go")

		Definition(false)

		MapsDeclarations(false)

		CommaOkIdiomInMaps(false)

		DeletingFromMaps(false)

		MapReadWrite(false)

		ComparingMaps(false)

		MapSet(false)
	}
}
