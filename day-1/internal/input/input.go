package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Extract() []int {
	file, err := ioutil.ReadFile("internal/input/input.txt")
	if err != nil {
		log.Fatal("failed to read file")
	}

	lines := strings.Split(string(file), "\n")
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
