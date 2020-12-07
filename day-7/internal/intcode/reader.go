package intcode

type Reader interface {
	RetrieveSignal() int
}

type MemReader struct {
	initialized  bool
	PhaseSetting int
	InputSignal  int
}

func (r *MemReader) RetrieveSignal() int {
	if !r.initialized {
		r.initialized = true
		return r.PhaseSetting
	}

	r.initialized = false
	return r.InputSignal
}

type ChanReader struct {
	initialized  bool
	PhaseSetting int
	InputSignal  <-chan int
}

func (r *ChanReader) RetrieveSignal() int {
	if !r.initialized {
		r.initialized = true
		return r.PhaseSetting
	}

	return <-r.InputSignal
}
