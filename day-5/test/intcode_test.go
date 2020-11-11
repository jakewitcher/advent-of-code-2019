package intcode

import (
	"day-5/internal/intcode"
	"testing"
)

var opCodeTests = []struct {
	program  []int
	modes    int
	index    int
	one, two int
}{
	{
		[]int{1101, 4, 5, 6, 10, 7, 0},
		11,
		0,
		9, 20,
	},
	{
		[]int{1, 4, 5, 6, 10, 7, 0},
		0,
		0,
		17, 70,
	},
	{
		[]int{1001, 4, 5, 6, 10, 7, 0},
		10,
		0,
		15, 50,
	},
	{
		[]int{101, 4, 5, 6, 10, 7, 0},
		1,
		0,
		11, 28,
	},
}

func TestProcessOpCodeOne(t *testing.T) {
	for _, test := range opCodeTests {
		intcode.ProcessOpCodeOne(test.program, test.modes, test.index)
		if test.program[6] != test.one {
			t.Fatalf("expected %v\ngot %v", test.one, test.program[6])
		}
	}
}

func TestProcessOpCodeTwo(t *testing.T) {
	for _, test := range opCodeTests {
		intcode.ProcessOpCodeTwo(test.program, test.modes, test.index)
		if test.program[6] != test.two {
			t.Fatalf("expected %v\ngot %v", test.two, test.program[6])
		}
	}
}
