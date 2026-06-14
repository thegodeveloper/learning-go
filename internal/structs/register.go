package structs

import "github.com/thegodeveloper/learning-go/internal/registry"

func init() {
	registry.Register(registry.NewModule("structs", Run, map[string]func(bool){
		"definition":      Definition,
		"demo":            Demo,
		"activity":        Activity,
		"records":         Records,
		"pointerReceiver": PointerReceiver,
		"methods":         Methods,
		"binaryTree":      BinaryTree,
		"songs":           Songs,
		"embedding":       Embedding,
	}))
}
