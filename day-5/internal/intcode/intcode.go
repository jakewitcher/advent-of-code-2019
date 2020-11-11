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
	Program []int
	pos     int
}

func (icc IntCodeComputer) Run() error {
	for icc.pos < len(icc.Program) {
		op, mode := icc.parseOpCode()

		switch op {
		case 1:
			icc.ProcessOpCodeOne(mode)
			icc.pos += 4

		case 2:
			icc.ProcessOpCodeTwo(mode)
			icc.pos += 4

		case 3:
			icc.ProcessOpCodeThree()
			icc.pos += 2

		case 4:
			icc.ProcessOpCodeFour(mode)
			icc.pos += 2

		case 99:
			return nil

		default:
			return errors.New(fmt.Sprintf("invalid op code %v", op))
		}
	}

	return errors.New("program exited unexpectedly")
}

func (icc IntCodeComputer) ProcessOpCodeOne(mode Mode) {
	p1, p2, p3 := icc.Program[icc.pos+1], icc.Program[icc.pos+2], icc.Program[icc.pos+3]

	l := icc.getOperand(mode.First(), p1)
	r := icc.getOperand(mode.Second(), p2)

	icc.Program[p3] = l + r
}

func (icc IntCodeComputer) ProcessOpCodeTwo(mode Mode) {
	p1, p2, p3 := icc.Program[icc.pos+1], icc.Program[icc.pos+2], icc.Program[icc.pos+3]

	l := icc.getOperand(mode.First(), p1)
	r := icc.getOperand(mode.Second(), p2)

	icc.Program[p3] = l * r
}

func (icc IntCodeComputer) ProcessOpCodeThree() {
	p := icc.Program[icc.pos+1]
	id := icc.GetSystemId()
	icc.Program[p] = id
}

func (icc IntCodeComputer) ProcessOpCodeFour(mode Mode) {
	p := icc.Program[icc.pos+1]
	o := icc.getOperand(mode.First(), p)
	icc.PrintCode(o)
}

func (icc IntCodeComputer) getOperand(mode int, p int) int {
	if mode == 1 {
		return p
	}

	return icc.Program[p]
}

func (icc IntCodeComputer) parseOpCode() (int, Mode) {
	val := icc.Program[icc.pos]
	op := val % 100
	mode := Mode{val / 100}

	return op, mode
}
