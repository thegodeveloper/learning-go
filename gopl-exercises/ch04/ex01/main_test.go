package main

import (
	"testing"
)

func TestSha256PopCount(t *testing.T) {
	var tests = []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"AIUEO", "AIUEO", 0},
		{"a", "b", 126},
		{"AIUEO", "Kakikukeko", 128},
	}

	for _, test := range tests {
		got := sha256PopCount([]byte(test.a), []byte(test.b))
		if got != test.want {
			t.Errorf("sha256PopCount(%s, %s) = %d, want %d", test.a, test.b, got, test.want)
		}
	}
}
