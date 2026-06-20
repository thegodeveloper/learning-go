package strings

import (
	"fmt"
	"strings"
)

func BuilderString(show bool) {
	if show {
		fmt.Println("-- String Builder")

		var builder strings.Builder

		// write some string
		builder.WriteString("Hello")
		builder.WriteString(", ")
		builder.WriteString("World")

		// convert builder to string
		result := builder.String()
		fmt.Println(result)
	}
}
