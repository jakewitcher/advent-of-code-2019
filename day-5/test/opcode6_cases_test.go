package intcode

import "day-5/internal/intcode"

var opCodeSixTests = []struct {
	process []int
	mode    intcode.Mode
	output  int
}{
	{
		[]int{1106, 1, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 11},
		3,
	},
	{
		[]int{1106, 0, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 11},
		5,
	},
	{
		[]int{6, 4, 5, 6, 1, 6, 0},
		intcode.Mode{Modes: 0},
		3,
	},
	{
		[]int{6, 4, 5, 6, 0, 6, 0},
		intcode.Mode{Modes: 0},
		6,
	},
	{
		[]int{1006, 4, 5, 6, 1, 7, 0},
		intcode.Mode{Modes: 10},
		3,
	},
	{
		[]int{1006, 4, 5, 6, 0, 7, 0},
		intcode.Mode{Modes: 10},
		5,
	},
	{
		[]int{106, 1, 5, 6, 6, 6, 0},
		intcode.Mode{Modes: 1},
		3,
	},
	{
		[]int{106, 0, 5, 6, 10, 6, 0},
		intcode.Mode{Modes: 1},
		6,
	},
}
