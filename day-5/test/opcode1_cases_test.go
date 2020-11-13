package intcode

import "day-5/internal/intcode"

var opCodeOneTests = []struct {
	process []int
	mode    intcode.Mode
	output  int
}{
	{
		[]int{1101, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 11},
		9,
	},
	{
		[]int{1, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 0},
		17,
	},
	{
		[]int{1001, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 10},
		15,
	},
	{
		[]int{101, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 1},
		11,
	},
}
