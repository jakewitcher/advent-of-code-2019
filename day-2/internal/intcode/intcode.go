package intcode

import (
	"errors"
	"fmt"
)

func Run(program []int, noun, verb int) ([]int, error) {
	program[1] = noun
	program[2] = verb

	for i := 0; i < len(program); i += 4 {
		opCode := program[i]

		switch opCode {
		case 1:
			l, r, t := program[i+1], program[i+2], program[i+3]
			program[t] = program[l] + program[r]

		case 2:
			l, r, t := program[i+1], program[i+2], program[i+3]
			program[t] = program[l] * program[r]

		case 99:
			return program, nil

		default:
			return nil, errors.New(fmt.Sprintf("invalid op code %v", i))
		}
	}

	return nil, errors.New("program exited unexpectedly")
}

func Produce(program []int, expect int) (int, error) {
	l := len(program)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			programCopy := make([]int, l)
			copy(programCopy, program)

			output, err := Run(programCopy, i, j)
			if err != nil {
				continue
			}

			if output[0] == 19690720 {
				return 100*i + j, nil
			}
		}
	}

	return 0, errors.New("no noun and verb produces expected output")
}
