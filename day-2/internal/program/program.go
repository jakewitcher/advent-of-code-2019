package program

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Extract() []int {
	file, err := ioutil.ReadFile("internal/program/program.txt")
	if err != nil {
		log.Fatal("failed to read file")
	}

	lines := strings.Split(string(file), ",")
	input := make([]int, len(lines))

	for i, line := range lines {
		n, err := strconv.Atoi(line)

		if err != nil {
			log.Fatalf("failed to parse %v as int", line)
		}

		input[i] = n
	}

	return input
}
