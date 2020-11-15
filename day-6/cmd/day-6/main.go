package main

import (
	"day-6/internal/input"
	"day-6/internal/orbits"
	"fmt"
)

func main() {
	in := input.Extract()
	PartOne(in)
	PartTwo(in)
}

func PartOne(input []string) {
	total := orbits.TotalNumberOfOrbits(input)
	fmt.Println(total)
}

func PartTwo(input []string) {
	transfers, err := orbits.OrbitalTransfers(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(transfers)
}
