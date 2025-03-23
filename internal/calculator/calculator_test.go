package calculator_test

import (
	"testing"

	"github.com/thegodeveloper/learning-go/internal/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	var want float64 = 49

	got := calculator.Add(7, 7)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
