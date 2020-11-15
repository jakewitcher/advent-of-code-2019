package intcode

type Mode struct {
	Modes int
}

func (m Mode) First() int {
	return m.Modes % 10
}

func (m Mode) Second() int {
	return (m.Modes / 10) % 10
}
