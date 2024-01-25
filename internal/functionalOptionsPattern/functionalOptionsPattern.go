package functionalOptionsPattern

import "fmt"

type Pizza struct {
	pepper         bool
	salt           bool
	cheese         bool
	extraCheese    bool
	chickenTopping bool
	paneerTopping  bool
}

// Options option type which return a function
type Options func(config *Pizza)

func BuildPizza(options ...Options) *Pizza {
	p := &Pizza{
		salt:   true,
		cheese: true,
	}

	for _, opt := range options {
		if opt != nil {
			opt(p)
		}
	}

	return p
}

func WIthPepper() func(config *Pizza) {
	return func(config *Pizza) {
		config.pepper = true
	}
}

func WithSalt() func(config *Pizza) {
	return func(config *Pizza) {
		config.salt = true
	}
}

func WithCheese() func(config *Pizza) {
	return func(config *Pizza) {
		config.cheese = true
	}
}

func WithExtraCheese() func(config *Pizza) {
	return func(config *Pizza) {
		config.extraCheese = true
	}
}

func WithChickenTopping() func(config *Pizza) {
	return func(config *Pizza) {
		config.chickenTopping = true
	}
}

func WithPaneerTopping() func(config *Pizza) {
	return func(config *Pizza) {
		config.paneerTopping = true
	}
}

func Master(show bool) {
	if show {
		coolPizza := BuildPizza(WIthPepper(), WithExtraCheese(), WithChickenTopping())

		fmt.Printf("eating delicious Pizza: %v", coolPizza)
	}
}
