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
			icc.moveForward(4)

		case 2:
			icc.ProcessOpCodeTwo(mode)
			icc.moveForward(4)

		case 3:
			icc.ProcessOpCodeThree()
			icc.moveForward(2)

		case 4:
			icc.ProcessOpCodeFour(mode)
			icc.moveForward(2)

		case 99:
			return nil

		default:
			return errors.New(fmt.Sprintf("invalid op code %v", op))
		}
	}

	return errors.New("program exited unexpectedly")
}

func (icc IntCodeComputer) ProcessOpCodeOne(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	i := icc.getIntAt(icc.pos + 3)

	icc.setIntAt(i, l+r)
}

func (icc IntCodeComputer) ProcessOpCodeTwo(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	icc.setIntAt(p, l*r)
}

func (icc IntCodeComputer) ProcessOpCodeThree() {
	p := icc.getIntAt(icc.pos + 1)
	id := icc.GetSystemId()
	icc.setIntAt(p, id)
}

func (icc IntCodeComputer) ProcessOpCodeFour(mode Mode) {
	p := icc.getIntAt(icc.pos + 1)
	o := icc.getOperand(mode.First(), p)
	icc.PrintCode(o)
}

func (icc IntCodeComputer) getLhsRhsOperands(mode Mode) (int, int) {
	p1, p2 := icc.getIntAt(icc.pos+1), icc.getIntAt(icc.pos+2)

	l := icc.getOperand(mode.First(), p1)
	r := icc.getOperand(mode.Second(), p2)

	return l, r
}

func (icc IntCodeComputer) getTargetIndex() int {
	return icc.getIntAt(icc.pos + 3)
}

func (icc IntCodeComputer) getOperand(mode int, p int) int {
	if mode == 1 {
		return p
	}

	return icc.Program[p]
}

func (icc IntCodeComputer) parseOpCode() (int, Mode) {
	n := icc.getIntAt(icc.pos)
	op := n % 100
	mode := Mode{n / 100}

	return op, mode
}

func (icc IntCodeComputer) getIntAt(n int) int {
	return icc.Program[n]
}

func (icc IntCodeComputer) setIntAt(i int, n int) {
	icc.Program[i] = n
}

func (icc *IntCodeComputer) moveForward(n int) {
	icc.pos += n
}
