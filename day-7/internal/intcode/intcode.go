package intcode

import (
	"log"
)

type Computer struct {
	Id string
	Reader
	Writer
	*Program
}

func (icc Computer) Run() {
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
			icc.Close(icc.Id)
			return

		default:
			log.Fatalf("invalid op code %v", op)
		}
	}

	log.Fatal("program exited unexpectedly")
}

func (icc Computer) ProcessOpCodeOne(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	icc.setIntAt(p, l+r)
	icc.moveForward(4)
}

func (icc Computer) ProcessOpCodeTwo(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	icc.setIntAt(p, l*r)
	icc.moveForward(4)
}

func (icc Computer) ProcessOpCodeThree() {
	p := icc.getIntAt(icc.pos + 1)
	id := icc.RetrieveSignal()

	icc.setIntAt(p, id)
	icc.moveForward(2)
}

func (icc Computer) ProcessOpCodeFour(mode Mode) {
	p := icc.getIntAt(icc.pos + 1)
	o := icc.getOperand(mode.First(), p)

	icc.SendSignal(icc.Id, o)
	icc.moveForward(2)
}

func (icc Computer) ProcessOpCodeFive(mode Mode) {
	p1, p2 := icc.getIntAt(icc.pos+1), icc.getIntAt(icc.pos+2)

	if icc.isTrue(mode.First(), p1) {
		p := icc.getOperand(mode.Second(), p2)
		icc.setPosition(p)
	} else {
		icc.moveForward(3)
	}
}

func (icc Computer) ProcessOpCodeSix(mode Mode) {
	p1, p2 := icc.getIntAt(icc.pos+1), icc.getIntAt(icc.pos+2)

	if !icc.isTrue(mode.First(), p1) {
		p := icc.getOperand(mode.Second(), p2)
		icc.setPosition(p)
	} else {
		icc.moveForward(3)
	}
}

func (icc Computer) ProcessOpCodeSeven(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	if l < r {
		icc.setIntAt(p, 1)
	} else {
		icc.setIntAt(p, 0)
	}

	icc.moveForward(4)
}

func (icc Computer) ProcessOpCodeEight(mode Mode) {
	l, r := icc.getLhsRhsOperands(mode)
	p := icc.getTargetIndex()

	if l == r {
		icc.setIntAt(p, 1)
	} else {
		icc.setIntAt(p, 0)
	}

	icc.moveForward(4)
}

func (icc Computer) parseOpCode() (int, Mode) {
	n := icc.getIntAt(icc.pos)
	op := n % 100
	mode := Mode{n / 100}

	return op, mode
}

func (icc Computer) getLhsRhsOperands(mode Mode) (int, int) {
	p1, p2 := icc.getIntAt(icc.pos+1), icc.getIntAt(icc.pos+2)

	l := icc.getOperand(mode.First(), p1)
	r := icc.getOperand(mode.Second(), p2)

	return l, r
}

func (icc Computer) getOperand(mode int, p int) int {
	if mode == 1 {
		return p
	}

	return icc.getIntAt(p)
}

func (icc Computer) getTargetIndex() int {
	return icc.getIntAt(icc.pos + 3)
}

func (icc Computer) isTrue(mode int, p int) bool {
	val := icc.getOperand(mode, p)
	return val != 0
}
