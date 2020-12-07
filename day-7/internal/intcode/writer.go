package intcode

type Writer interface {
	SendSignal(source string, value int)
	Close(source string)
}

type MemWriter struct {
	OutputSignal int
}

func (w *MemWriter) SendSignal(source string, value int) {
	w.OutputSignal = value
}

func (w *MemWriter) Close(source string) {

}

type ChanWriter struct {
	OutputSignal chan<- Signal
}

func (w *ChanWriter) SendSignal(source string, value int) {
	w.OutputSignal <- Signal{Source: source, Value: value}
}

func (w *ChanWriter) Close(source string) {
	w.SendSignal(source, -1)
}
