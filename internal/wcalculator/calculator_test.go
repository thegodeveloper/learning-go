package wcalculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 7, b: 0, want: 7},
	}

	for _, tc := range testCases {
		got := Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 1, b: 1, want: 0},
		{a: 7, b: -14, want: 21},
		{a: 0, b: 7, want: -7},
	}

	for _, tc := range testCases {
		got := Subtract(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 4, b: 2, want: 8},
		{a: 1, b: 1, want: 1},
		{a: 7, b: 11, want: 77},
	}

	for _, tc := range testCases {
		got := Multiply(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}

	for _, tc := range testCases {
		got, err := Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
