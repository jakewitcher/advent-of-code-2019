package main

import (
	"day-4/internal/input"
	"day-4/internal/passwords"
	"fmt"
)

func main() {
	min, max := input.Extract()
	count := passwords.ValidPasswords(min, max)
	fmt.Println(count)
}
