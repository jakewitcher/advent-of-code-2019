package fuel

import (
	"day-1/internal/fuel"
	"testing"
)

var testCases = []struct {
	input, simple, recursive int
}{
	{12, 2, 2},
	{14, 2, 2},
	{1969, 654, 966},
	{100756, 33583, 50346},
}

func TestCalculate(t *testing.T) {
	for _, test := range testCases {
		if fuel := fuel.Calculate(test.input); fuel != test.simple {
			t.Fatalf("got %v\nexpected %v", fuel, test.simple)
		}
	}
}

func TestCalculateAll(t *testing.T) {
	for _, test := range testCases {
		if fuel := fuel.CalculateAll(test.input); fuel != test.recursive {
			t.Fatalf("got %v\nexpected %v", fuel, test.recursive)
		}
	}
}
