package amplifier

import (
	"day-7/internal/intcode"
	"testing"
)

var powerThrusterCases = []struct {
	sequence []int
	process  []int
	output   int
}{
	{
		[]int{4, 3, 2, 1, 0},
		[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 4, 0},
		43210,
	},
	{
		[]int{0, 1, 2, 3, 4},
		[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		54321,
	},
	{
		[]int{1, 0, 4, 3, 2},
		[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
		65210,
	},
}

func TestPowerThrusters(t *testing.T) {
	for _, test := range powerThrusterCases {
		output := intcode.PowerThrusters(test.process, test.sequence)
		if output != test.output {
			t.Fatalf("expected %v, got %v", test.output, output)
		}
	}
}

