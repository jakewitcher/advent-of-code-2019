package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Extract() (int, int) {
	file, err := ioutil.ReadFile("internal/input/input.txt")
	if err != nil {
		log.Fatal("failed to read file")
	}

	vals := strings.Split(string(file), "-")

	min, err := strconv.Atoi(vals[0])

	if err != nil {
		log.Fatalf("invalid min value %v", vals[0])
	}

	max, err := strconv.Atoi(vals[1])

	if err != nil {
		log.Fatalf("invalid max value %v", vals[1])
	}

	return min, max
}
