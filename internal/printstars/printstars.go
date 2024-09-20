package printstars

import (
	"fmt"
	"math/rand"
	"strings"
)

func Master(show bool) {
	if show {
		fmt.Println("---Print Stars---")
		r := rand.Intn(7) + 1
		stars := strings.Repeat("*", r)
		fmt.Println(stars)
	}
}
