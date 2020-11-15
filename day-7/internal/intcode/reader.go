package intcode

type Reader interface {
	GetNext() int
}

type MemReader struct{
	count int
	PhaseSetting int
	InputSignal int
}

func (r *MemReader) GetNext() int {
	if r.count == 0 {
		r.count++
		return r.PhaseSetting
	}

	return r.InputSignal
}
