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

func Generate(start, end int) [][]int {
	var sequences [][]int

	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			for k := start; k < end; k++ {
				for l := start; l < end; l++ {
					for m := start; m < end; m++ {
						if i != j && i != k && i != l && i != m &&
							j != k && j != l && j != m && k != l && k != m && l != m {
							sequences = append(sequences, []int{i, j, k, l, m})
						}
					}
				}
			}
		}
	}

	return sequences
}
