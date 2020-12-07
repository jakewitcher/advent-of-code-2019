package main

import (
	"day-7/internal/input"
	"day-7/internal/intcode"
	"fmt"
)

func main() {
	process := input.Extract()
	PartOne(process)
	PartTwo(process)
}

func PartOne(process []int) {
	sequences := input.Generate(0, 5)
	var max int

	for _, sequence := range sequences {
		signal := intcode.PowerThrusters(process, sequence)

		if signal > max {
			max = signal
		}
	}

	fmt.Println(max)
}

func PartTwo(process []int) {
	sequences := input.Generate(5, 10)
	var max int

	for _, sequence := range sequences {
		signal := intcode.FeedbackLoop(process, sequence)

		if signal > max {
			max = signal
		}
	}

	fmt.Println(max)
}
