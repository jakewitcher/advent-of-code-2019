package input

import (
	"io/ioutil"
	"log"
	"strings"
)

func Extract() ([]string, []string) {
	file, err := ioutil.ReadFile("internal/input/input.txt")
	if err != nil {
		log.Fatal("failed to read file")
	}

	lines := strings.Split(string(file), "\n")

	if len(lines) != 2 {
		log.Fatal("number of wires must equal 2")
	}

	wireA, wireB := strings.Split(lines[0], ","), strings.Split(lines[1], ",")

	return wireA, wireB
}
