package output

import "fmt"

type Writer interface {
	PrintCode(n int)
}

type IntCodeWriter struct{}

func (IntCodeWriter) PrintCode(code int) {
	fmt.Println(code)
}
