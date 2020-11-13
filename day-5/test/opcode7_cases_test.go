package intcode

import "day-5/internal/intcode"

var opCodeSevenTests = []struct {
	process []int
	mode    intcode.Mode
	output  int
}{
	{
		[]int{1107, 1, 5, 6, 10, 7, 9},
		intcode.Mode{Modes: 11},
		1,
	},
	{
		[]int{1107, 8, 5, 6, 10, 7, 9},
		intcode.Mode{Modes: 11},
		0,
	},
	{
		[]int{7, 4, 5, 6, 1, 6, 9},
		intcode.Mode{Modes: 0},
		1,
	},
	{
		[]int{7, 4, 5, 6, 7, 3, 9},
		intcode.Mode{Modes: 0},
		0,
	},
	{
		[]int{1007, 4, 5, 6, 1, 7, 9},
		intcode.Mode{Modes: 10},
		1,
	},
	{
		[]int{1007, 4, 5, 6, 8, 7, 9},
		intcode.Mode{Modes: 10},
		0,
	},
	{
		[]int{107, 4, 5, 6, 3, 6, 9},
		intcode.Mode{Modes: 1},
		1,
	},
	{
		[]int{107, 4, 5, 6, 8, 3, 9},
		intcode.Mode{Modes: 1},
		0,
	},
}
