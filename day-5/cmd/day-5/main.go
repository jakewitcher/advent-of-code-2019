package main

import (
	"day-5/internal/input"
	"day-5/internal/intcode"
	"day-5/internal/output"
	"fmt"
)

func main() {
	process := input.Extract()
	intCodeComputer := intcode.IntCodeComputer{
		Reader:  input.IntCodeReader{},
		Writer:  output.IntCodeWriter{},
		Program: &intcode.Program{Process: process},
	}

	err := intCodeComputer.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
