package calculator_test

import (
	"testing"

	"github.com/thegodeveloper/learning-go/internal/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	var want float64 = 14

	got := calculator.Add(7, 7)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	var want float64 = 7

	got := calculator.Subtract(14, 7)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
