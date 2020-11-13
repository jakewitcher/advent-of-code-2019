package intcode

import "day-5/internal/intcode"

var opCodeFourTests = []struct {
	process []int
	mode    intcode.Mode
	output  int
}{
	{
		[]int{14, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 1},
		4,
	},
	{
		[]int{14, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 0},
		10,
	},
}
