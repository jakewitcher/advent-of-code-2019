package input

import (
	"io/ioutil"
	"log"
	"strings"
)

func Extract() []string {
	file, err := ioutil.ReadFile("internal/input/input.txt")
	if err != nil {
		log.Fatal("failed to read file")
	}

	orbits := strings.Split(string(file), "\r\n")

	return orbits
}
