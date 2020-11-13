package intcode

type Program struct {
	Process []int
	pos     int
}

func (p Program) getIntAt(n int) int {
	return p.Process[n]
}

func (p Program) setIntAt(i int, n int) {
	p.Process[i] = n
}

func (p *Program) moveForward(n int) {
	p.pos += n
}

func (p *Program) setPosition(n int) {
	p.pos = n
}

func (p Program) Processing() bool {
	return p.pos < len(p.Process)
}

func (p Program) CurrentPosition() int {
	return p.pos
}
