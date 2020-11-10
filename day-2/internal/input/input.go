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

	vals := strings.Split(string(file), ",")
	input := make([]int, len(vals))

	for i, val := range vals {
		n, err := strconv.Atoi(val)

		if err != nil {
			log.Fatalf("failed to parse %v as int", val)
		}

		input[i] = n
	}

	return input
}
