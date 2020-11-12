package intcode

import (
	"day-5/internal/intcode"
	"testing"
)

var opCodeTests = []struct {
	program  []int
	mode     intcode.Mode
	one, two int
}{
	{
		[]int{1101, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 11},
		9, 20,
	},
	{
		[]int{1, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 0},
		17, 70,
	},
	{
		[]int{1001, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 10},
		15, 50,
	},
	{
		[]int{101, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 1},
		11, 28,
	},
}

func TestProcessOpCodeOne(t *testing.T) {
	for _, test := range opCodeTests {
		intcode.IntCodeComputer{Program: test.program}.ProcessOpCodeOne(test.mode)
		if test.program[6] != test.one {
			t.Fatalf("expected %v\ngot %v", test.one, test.program[6])
		}
	}
}

func TestProcessOpCodeTwo(t *testing.T) {
	for _, test := range opCodeTests {
		intcode.IntCodeComputer{Program: test.program}.ProcessOpCodeTwo(test.mode)
		if test.program[6] != test.two {
			t.Fatalf("expected %v\ngot %v", test.two, test.program[6])
		}
	}
}
