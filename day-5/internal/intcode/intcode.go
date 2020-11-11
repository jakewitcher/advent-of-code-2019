package intcode

import (
	"day-5/internal/input"
	"day-5/internal/output"
	"errors"
	"fmt"
)

func Run(program []int, reader input.Reader, writer output.Writer) error {
	for i := 0; i < len(program); {
		opCode := program[i]
		op := opCode % 100
		modes := opCode / 100

		switch op {
		case 1:
			ProcessOpCodeOne(program, modes, i)
			i += 4

		case 2:
			ProcessOpCodeTwo(program, modes, i)
			i += 4

		case 3:
			ProcessOpCodeThree(program, i, reader)
			i += 2

		case 4:
			ProcessOpCodeFour(program, modes, i, writer)
			i += 2

		case 99:
			return nil

		default:
			return errors.New(fmt.Sprintf("invalid op code %v", i))
		}
	}

	return errors.New("program exited unexpectedly")
}

func ProcessOpCodeOne(program []int, modes int, i int) {
	p1, p2, p3 := program[i+1], program[i+2], program[i+3]

	l := getOperand(program, modes%10, p1)
	r := getOperand(program, (modes/10)%10, p2)

	program[p3] = l + r
}

func ProcessOpCodeTwo(program []int, modes int, i int) {
	p1, p2, p3 := program[i+1], program[i+2], program[i+3]

	l := getOperand(program, modes%10, p1)
	r := getOperand(program, (modes/10)%10, p2)

	program[p3] = l * r
}

func ProcessOpCodeThree(program []int, i int, reader input.Reader) {
	p := program[i+1]
	id := reader.GetSystemId()
	program[p] = id
}

func ProcessOpCodeFour(program []int, modes int, i int, writer output.Writer) {
	p := program[i+1]
	o := getOperand(program, modes%10, p)
	writer.PrintCode(o)
}

func getOperand(program []int, mode int, p int) int {
	if mode == 1 {
		return p
	}

	return program[p]
}
