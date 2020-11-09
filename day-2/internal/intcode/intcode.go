package intcode

import (
	"errors"
	"fmt"
)

func Run(prog []int, noun, verb int) ([]int, error) {
	prog[1] = noun
	prog[2] = verb

	for i := 0; i < len(prog); i += 4 {
		opCode := prog[i]

		switch opCode {
		case 1:
			l, r, t := prog[i+1], prog[i+2], prog[i+3]
			prog[t] = prog[l] + prog[r]

		case 2:
			l, r, t := prog[i+1], prog[i+2], prog[i+3]
			prog[t] = prog[l] * prog[r]

		case 99:
			return prog, nil

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
			input := make([]int, l)
			copy(input, program)

			output, err := Run(input, i, j)
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
