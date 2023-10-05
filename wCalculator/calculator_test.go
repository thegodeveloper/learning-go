package wCalculator_test

import (
	"github.com/thegodeveloper/learning-go/wCalculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4

	got := wCalculator.Add(2, 2)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
