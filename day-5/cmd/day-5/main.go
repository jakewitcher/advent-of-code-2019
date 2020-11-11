package main

import (
	"day-5/internal/input"
	"day-5/internal/intcode"
	"day-5/internal/output"
	"fmt"
)

func main() {
	program := input.Extract()
	err := intcode.Run(program, input.IntCodeReader{}, output.IntCodeWriter{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
