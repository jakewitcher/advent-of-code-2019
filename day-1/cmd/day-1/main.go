package main

import (
	"day-1/internal/fuel"
	"day-1/internal/input"
	"fmt"
)

func main() {
	input := input.Extract()
	PartOne(input)
	PartTwo(input)
}

func PartOne(input []int) {
	var sum int
	for _, mass := range input {
		sum += fuel.Calculate(mass)
	}

	fmt.Println(sum)
}

func PartTwo(input []int) {
	var sum int
	for _, mass := range input {
		sum += fuel.CalculateAll(mass)
	}

	fmt.Println(sum)
}
