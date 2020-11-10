package main

import (
	"day-3/internal/input"
	"day-3/internal/wires"
	"fmt"
)

func main() {
	wireA, wireB := input.Extract()
	PartOne(wireA, wireB)
	PartTwo(wireA, wireB)
}

func PartOne(wireA, wireB []string) {
	distance := wires.ClosestIntersection(wireA, wireB)
	fmt.Println(distance)
}

func PartTwo(wireA, wireB []string) {
	steps := wires.ShortestSteps(wireA, wireB)
	fmt.Println(steps)
}
