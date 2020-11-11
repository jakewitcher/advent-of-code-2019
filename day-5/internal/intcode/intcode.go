package intcode

import (
	"day-5/internal/input"
	"day-5/internal/output"
	"errors"
	"fmt"
)

type IntCodeComputer struct {
	input.Reader
	output.Writer
}

func (icc IntCodeComputer) Run(program []int) error {
	for i := 0; i < len(program); {
		opCode := program[i]
		op := opCode % 100
		modes := opCode / 100

		switch op {
		case 1:
			icc.ProcessOpCodeOne(program, modes, i)
			i += 4

		case 2:
			icc.ProcessOpCodeTwo(program, modes, i)
			i += 4

		case 3:
			icc.ProcessOpCodeThree(program, i)
			i += 2

		case 4:
			icc.ProcessOpCodeFour(program, modes, i)
			i += 2

		case 99:
			return nil

		default:
			return errors.New(fmt.Sprintf("invalid op code %v", i))
		}
	}

	return errors.New("program exited unexpectedly")
}

func (IntCodeComputer) ProcessOpCodeOne(program []int, modes int, i int) {
	p1, p2, p3 := program[i+1], program[i+2], program[i+3]

	l := getOperand(program, modes%10, p1)
	r := getOperand(program, (modes/10)%10, p2)

	program[p3] = l + r
}

func (IntCodeComputer) ProcessOpCodeTwo(program []int, modes int, i int) {
	p1, p2, p3 := program[i+1], program[i+2], program[i+3]

	l := getOperand(program, modes%10, p1)
	r := getOperand(program, (modes/10)%10, p2)

	program[p3] = l * r
}

func (icc IntCodeComputer) ProcessOpCodeThree(program []int, i int) {
	p := program[i+1]
	id := icc.GetSystemId()
	program[p] = id
}

func (icc IntCodeComputer) ProcessOpCodeFour(program []int, modes int, i int) {
	p := program[i+1]
	o := getOperand(program, modes%10, p)
	icc.PrintCode(o)
}

func getOperand(program []int, mode int, p int) int {
	if mode == 1 {
		return p
	}

	return program[p]
}
