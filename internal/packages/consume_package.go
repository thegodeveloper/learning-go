package packages

import (
	"fmt"

	"github.com/thegodeveloper/learning-go/internal/functions"
)

func ConsumePackage(show bool) {
	if show {
		fmt.Println("--- Consume Package ---")

		functions.Message("Bill")
	}
}
