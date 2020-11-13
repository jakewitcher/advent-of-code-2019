package intcode

import "day-5/internal/intcode"

var opCodeEightTests = []struct {
	process []int
	mode    intcode.Mode
	output  int
}{
	{
		[]int{1108, 1, 1, 6, 10, 7, 9},
		intcode.Mode{Modes: 11},
		1,
	},
	{
		[]int{1108, 8, 5, 6, 10, 7, 9},
		intcode.Mode{Modes: 11},
		0,
	},
	{
		[]int{8, 4, 5, 6, 3, 3, 9},
		intcode.Mode{Modes: 0},
		1,
	},
	{
		[]int{8, 4, 5, 6, 7, 3, 9},
		intcode.Mode{Modes: 0},
		0,
	},
	{
		[]int{1008, 4, 5, 6, 5, 7, 9},
		intcode.Mode{Modes: 10},
		1,
	},
	{
		[]int{1008, 4, 5, 6, 8, 7, 9},
		intcode.Mode{Modes: 10},
		0,
	},
	{
		[]int{108, 4, 5, 6, 3, 4, 9},
		intcode.Mode{Modes: 1},
		1,
	},
	{
		[]int{108, 4, 5, 6, 8, 3, 9},
		intcode.Mode{Modes: 1},
		0,
	},
}
