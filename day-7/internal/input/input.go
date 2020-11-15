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

func Generate() [][]int {
	var sequences [][]int

	for i := 0; i < 5; i ++ {
		for j := 0; j < 5; j ++ {
			for k := 0; k < 5; k ++ {
				for l := 0; l < 5; l ++ {
					for m := 0; m < 5; m ++ {
						if i != j && i != k && i != l && i != m &&
							j != k && j != l && j != m && k != l && k != m && l != m {
							sequences = append(sequences, []int{i,j,k,l,m})
						}
					}
				}
			}
		}
	}

	return sequences
}
