package input

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Reader interface {
	GetSystemId() int
}

type IntCodeReader struct{}

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

func (IntCodeReader) GetSystemId() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please provide system ID:")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("failed to get system id from user")
	}

	id, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		log.Fatalf("%v is not a valid system id", text)
	}

	return id
}
