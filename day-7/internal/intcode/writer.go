package intcode

type Writer interface {
	WriteSignal(n int)
}

type MemWriter struct{
	OutputSignal int
}

func (w *MemWriter) WriteSignal(code int) {
	w.OutputSignal = code
}
