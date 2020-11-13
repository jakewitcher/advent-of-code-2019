package intcode

import "day-5/internal/intcode"

var opCodeFiveTests = []struct {
	process []int
	mode    intcode.Mode
	output  int
}{
	{
		[]int{1105, 1, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 11},
		5,
	},
	{
		[]int{1105, 0, 5, 6, 10, 7, 0},
		intcode.Mode{Modes: 11},
		3,
	},
	{
		[]int{5, 4, 5, 6, 1, 6, 0},
		intcode.Mode{Modes: 0},
		6,
	},
	{
		[]int{5, 4, 5, 6, 0, 6, 0},
		intcode.Mode{Modes: 0},
		3,
	},
	{
		[]int{1005, 4, 5, 6, 1, 7, 0},
		intcode.Mode{Modes: 10},
		5,
	},
	{
		[]int{1005, 4, 5, 6, 0, 7, 0},
		intcode.Mode{Modes: 10},
		3,
	},
	{
		[]int{105, 1, 5, 6, 6, 6, 0},
		intcode.Mode{Modes: 1},
		6,
	},
	{
		[]int{105, 0, 5, 6, 10, 6, 0},
		intcode.Mode{Modes: 1},
		3,
	},
}
