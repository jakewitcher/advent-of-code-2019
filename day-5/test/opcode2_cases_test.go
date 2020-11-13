package intcode

import "day-5/internal/intcode"

var opCodeTwoTests = []struct {
	process []int
	mode    intcode.Mode
	output  int
}{
	{
		[]int{1102, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 11},
		20,
	},
	{
		[]int{2, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 0},
		70,
	},
	{
		[]int{1002, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 10},
		50,
	},
	{
		[]int{102, 4, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 1},
		28,
	},
}
