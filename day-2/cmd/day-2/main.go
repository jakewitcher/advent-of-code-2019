package main

import (
	"day-2/internal/input"
	"day-2/internal/intcode"
	"fmt"
	"log"
)

func main() {
	program := input.Extract()
	PartOne(program)

	program = input.Extract()
	PartTwo(program)
}

func PartOne(program []int) {
	output, err := intcode.Run(program, 12, 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output[0])
}

func PartTwo(program []int) {
	output, err := intcode.Produce(program, 19690720)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
