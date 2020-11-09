package main

import (
	"day-2/internal/intcode"
	"day-2/internal/program"
	"fmt"
	"log"
)

func main() {
	prog := program.Extract()
	PartOne(prog)

	prog = program.Extract()
	PartTwo(prog)
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
