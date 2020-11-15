package main

import (
	"day-7/internal/input"
	"day-7/internal/intcode"
	"fmt"
)

func main() {
	process := input.Extract()
	sequences := input.Generate()
	PartOne(process, sequences)
}

func PartOne(process []int, sequences [][]int) {
	var max int

	for _, sequence := range sequences {
		signal, err := intcode.PowerThrusters(process, sequence)

		if err != nil {
			fmt.Println(err)
			return
		}

		if signal > max {
			max = signal
		}
	}

	fmt.Println(max)
}