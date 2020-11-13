package intcode

import (
	"day-5/internal/intcode"
	"testing"
)

type IntCodeTestReader struct {
	id int
}

type IntCodeTestWriter struct {
	output int
}

func (r IntCodeTestReader) GetSystemId() int {
	return r.id
}

func (w *IntCodeTestWriter) PrintCode(code int) {
	w.output = code
}

func TestProcessOpCodeOne(t *testing.T) {
	for _, test := range opCodeOneTests {
		icc := intcode.IntCodeComputer{Program: &intcode.Program{Process: test.process}}
		icc.ProcessOpCodeOne(test.mode)

		if test.process[6] != test.output {
			t.Fatalf("expected %v, got %v", test.output, test.process[6])
		}
	}
}

func TestProcessOpCodeTwo(t *testing.T) {
	for _, test := range opCodeTwoTests {
		icc := intcode.IntCodeComputer{Program: &intcode.Program{Process: test.process}}
		icc.ProcessOpCodeTwo(test.mode)

		if test.process[6] != test.output {
			t.Fatalf("expected %v, got %v", test.output, test.process[6])
		}
	}
}

func TestProcessOpCodeThree(t *testing.T) {
	for _, test := range opCodeThreeTests {
		icc := intcode.IntCodeComputer{
			Reader:  IntCodeTestReader{id: test.id},
			Program: &intcode.Program{Process: test.process},
		}
		icc.ProcessOpCodeThree()

		if test.process[test.output] != test.id {
			t.Fatalf("expected %v, got %v", test.output, test.process[6])
		}
	}
}

func TestProcessOpCodeFour(t *testing.T) {
	for _, test := range opCodeFourTests {
		writer := IntCodeTestWriter{}

		icc := intcode.IntCodeComputer{
			Writer:  &writer,
			Program: &intcode.Program{Process: test.process},
		}
		icc.ProcessOpCodeFour(test.mode)

		if writer.output != test.output {
			t.Fatalf("expected %v, got %v", test.output, test.process[6])
		}
	}
}

func TestProcessOpCodeFive(t *testing.T) {
	for _, test := range opCodeFiveTests {
		icc := intcode.IntCodeComputer{Program: &intcode.Program{Process: test.process}}
		icc.ProcessOpCodeFive(test.mode)

		if icc.CurrentPosition() != test.output {
			t.Fatalf("expected %v, got %v", test.output, icc.CurrentPosition())
		}
	}
}

func TestProcessOpCodeSix(t *testing.T) {
	for _, test := range opCodeSixTests {
		icc := intcode.IntCodeComputer{Program: &intcode.Program{Process: test.process}}
		icc.ProcessOpCodeSix(test.mode)

		if icc.CurrentPosition() != test.output {
			t.Fatalf("expected %v, got %v", test.output, icc.CurrentPosition())
		}
	}
}

func TestProcessOpCodeSeven(t *testing.T) {
	for _, test := range opCodeSevenTests {
		icc := intcode.IntCodeComputer{Program: &intcode.Program{Process: test.process}}
		icc.ProcessOpCodeSeven(test.mode)

		if test.process[6] != test.output {
			t.Fatalf("expected %v, got %v", test.output, test.process[6])
		}
	}
}

func TestProcessOpCodeEight(t *testing.T) {
	for _, test := range opCodeEightTests {
		icc := intcode.IntCodeComputer{Program: &intcode.Program{Process: test.process}}
		icc.ProcessOpCodeEight(test.mode)

		if test.process[6] != test.output {
			t.Fatalf("expected %v, got %v", test.output, test.process[6])
		}
	}
}
