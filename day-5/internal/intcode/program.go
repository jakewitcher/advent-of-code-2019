package intcode

type Program struct {
	Sequence []int
	pos      int
}

func (p Program) getIntAt(n int) int {
	return p.Sequence[n]
}

func (p Program) setIntAt(i int, n int) {
	p.Sequence[i] = n
}

func (p *Program) moveForward(n int) {
	p.pos += n
}

func (p *Program) setPosition(n int) {
	p.pos = n
}

func (p Program) Processing() bool {
	return p.pos < len(p.Sequence)
}
