package wexternalpackages

import (
	"fmt"
	"log"

	"github.com/learning-go-book/formatter"
	"github.com/shopspring/decimal"
)

func Index(show bool) {
	if show {
		fmt.Println("-- Working with External Packages")
		definition("99.99", "7.25")
	}
}

func definition(a string, p string) {
	amount, err := decimal.NewFromString(a)
	if err != nil {
		log.Fatal(err)
	}
	percent, err := decimal.NewFromString(p)
	if err != nil {
		log.Fatal(err)
	}
	percent = percent.Div(decimal.NewFromInt(100))
	total := amount.Add(amount.Mul(percent)).Round(2)
	fmt.Println(formatter.Space(80, a, p, total.StringFixed(2)))
}
