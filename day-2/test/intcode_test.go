package intcode

import (
	"day-2/internal/intcode"
	"testing"
)

var testCases = []struct {
	noun, verb int
	program    []int
	output     []int
}{
	{
		0, 0,
		[]int{1, 0, 0, 0, 99},
		[]int{2, 0, 0, 0, 99},
	},
	{
		3, 0,
		[]int{2, 0, 0, 3, 99},
		[]int{2, 3, 0, 6, 99},
	},
	{
		4, 4,
		[]int{2, 0, 0, 5, 99, 0},
		[]int{2, 4, 4, 5, 99, 9801},
	},
	{
		1, 1,
		[]int{1, 0, 0, 4, 99, 5, 6, 0, 99},
		[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
	{
		9, 10,
		[]int{1, 0, 0, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
	},
}

func TestRun(t *testing.T) {
	for _, test := range testCases {
		if output, _ := intcode.Run(test.program, test.noun, test.verb); !isEqual(output, test.output) {
			t.Fatal("actual output does not match expected output for test")
		}
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
