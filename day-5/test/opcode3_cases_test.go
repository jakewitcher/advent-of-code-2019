package intcode

var opCodeThreeTests = []struct {
	process []int
	id      int
	output  int
}{
	{
		[]int{3, 6, 5, 6, 10, 7, 0},
		7,
		6,
	},
	{
		[]int{3, 5, 5, 6, 10, 7, 0},
		3,
		5,
	},
}
