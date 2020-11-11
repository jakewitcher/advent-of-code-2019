package intcode

type Mode struct {
	modes int
}

func (m Mode) First() int {
	return m.modes % 10
}

func (m Mode) Second() int {
	return (m.modes / 10) % 10
}
