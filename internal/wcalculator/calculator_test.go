package wcalculator_test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4

	got := wcalculator.Add(2, 2)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
