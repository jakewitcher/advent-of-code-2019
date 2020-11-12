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
	*Program
}

func (icc IntCodeComputer) Run() error {
	for icc.Processing() {
		op, mode := icc.parseOpCode()

		switch op {
		case 1:
			icc.ProcessOpCodeOne(mode)

		case 2:
			icc.ProcessOpCodeTwo(mode)

		case 3:
			icc.ProcessOpCodeThree()

		case 4:
			icc.ProcessOpCodeFour(mode)

		case 5:
			icc.ProcessOpCodeFive(mode)

		case 6:
			icc.ProcessOpCodeSix(mode)

		case 7:
			icc.ProcessOpCodeSeven(mode)

		case 8:
			icc.ProcessOpCodeEight(mode)

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
	p := icc.getTargetIndex()

	icc.setIntAt(p, l+r)
	icc.moveForward(4)
}

func (icc IntCodeComputer) ProcessOpCodeTwo(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	icc.setIntAt(p, l*r)
	icc.moveForward(4)
}

func (icc IntCodeComputer) ProcessOpCodeThree() {
	p := icc.getIntAt(icc.pos + 1)
	id := icc.GetSystemId()

	icc.setIntAt(p, id)
	icc.moveForward(2)
}

func (icc IntCodeComputer) ProcessOpCodeFour(mode Mode) {
	p := icc.getIntAt(icc.pos + 1)
	o := icc.getOperand(mode.First(), p)

	icc.PrintCode(o)
	icc.moveForward(2)
}

func (icc IntCodeComputer) ProcessOpCodeFive(mode Mode) {
	p1, p2 := icc.getIntAt(icc.pos+1), icc.getIntAt(icc.pos+2)

	if icc.isTrue(mode.First(), p1) {
		p := icc.getOperand(mode.Second(), p2)
		icc.setPosition(p)
	} else {
		icc.moveForward(3)
	}
}

func (icc IntCodeComputer) ProcessOpCodeSix(mode Mode) {
	p1, p2 := icc.getIntAt(icc.pos+1), icc.getIntAt(icc.pos+2)

	if !icc.isTrue(mode.First(), p1) {
		p := icc.getOperand(mode.Second(), p2)
		icc.setPosition(p)
	} else {
		icc.moveForward(3)
	}
}

func (icc IntCodeComputer) ProcessOpCodeSeven(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	if l < r {
		icc.setIntAt(p, 1)
	} else {
		icc.setIntAt(p, 0)
	}

	icc.moveForward(4)
}

func (icc IntCodeComputer) ProcessOpCodeEight(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	if l == r {
		icc.setIntAt(p, 1)
	} else {
		icc.setIntAt(p, 0)
	}

	icc.moveForward(4)
}

func (icc IntCodeComputer) parseOpCode() (int, Mode) {
	n := icc.getIntAt(icc.pos)
	op := n % 100
	mode := Mode{n / 100}

	return op, mode
}

func (icc IntCodeComputer) getLhsRhsOperands(mode Mode) (int, int) {
	p1, p2 := icc.getIntAt(icc.pos+1), icc.getIntAt(icc.pos+2)

	l := icc.getOperand(mode.First(), p1)
	r := icc.getOperand(mode.Second(), p2)

	return l, r
}

func (icc IntCodeComputer) getOperand(mode int, p int) int {
	if mode == 1 {
		return p
	}

	return icc.getIntAt(p)
}

func (icc IntCodeComputer) getTargetIndex() int {
	return icc.getIntAt(icc.pos + 3)
}

func (icc IntCodeComputer) isTrue(mode int, p int) bool {
	val := icc.getOperand(mode, p)
	return val != 0
}
