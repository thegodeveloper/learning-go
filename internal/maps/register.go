package maps

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("maps", Run, map[string]func(bool){
		"introduction": Introduction,
		"definition":   Definition,
		"declarations": MapsDeclarations,
		"commaOkIdiom": CommaOkIdiomInMaps,
		"deleting":     DeletingFromMaps,
		"readWrite":    MapReadWrite,
		"comparing":    ComparingMaps,
		"set":          MapSet,
		"mapLiteral":   MapLiteral,
		"mapNil":       MapNil,
	}))
}
