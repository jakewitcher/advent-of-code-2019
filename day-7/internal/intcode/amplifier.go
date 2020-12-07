package intcode

type Orchestrator struct {
	Amplifiers []Computer
	Receiver   chan Signal
	Senders    map[string]chan int
}

type Signal struct {
	Source string
	Value  int
}

func PowerThrusters(process, sequence []int) int {
	writer := MemWriter{}

	for _, phase := range sequence {
		p := make([]int, len(process))
		copy(p, process)

		reader := MemReader{
			PhaseSetting: phase,
			InputSignal:  writer.OutputSignal,
		}

		intCodeComputer := Computer{
			Reader:  &reader,
			Writer:  &writer,
			Program: &Program{Process: p},
		}

		intCodeComputer.Run()
	}

	return writer.OutputSignal
}

func FeedbackLoop(process, sequence []int) int {
	ids := []string{"A", "B", "C", "D", "E"}
	senders := make(map[string]chan int)

	orchestrator := createOrchestrator(senders)
	initializeAmplifiers(process, sequence, ids, orchestrator)

	for _, amplifier := range orchestrator.Amplifiers {
		go amplifier.Run()
	}

	orchestrator.Senders["A"] <- 0

	var value int
	for signal := range orchestrator.Receiver {
		target := getSignalTarget(signal)

		if signal.Value == -1 {
			close(orchestrator.Senders[signal.Source])
			delete(orchestrator.Senders, signal.Source)

			if len(orchestrator.Senders) == 0 {
				break
			}

			continue
		}

		value = signal.Value
		if _, ok := orchestrator.Senders[target]; ok {
			orchestrator.Senders[target] <- signal.Value
		} else {
			break
		}
	}

	return value
}

func createOrchestrator(senders map[string]chan int) Orchestrator {
	orchestrator := Orchestrator{
		Amplifiers: make([]Computer, 5),
		Receiver:   make(chan Signal),
		Senders:    senders,
	}
	return orchestrator
}

func initializeAmplifiers(process []int, sequence []int, ids []string, orchestrator Orchestrator) {
	for i, id := range ids {
		orchestrator.Senders[id] = make(chan int)

		computer := Computer{
			Id: id,
			Reader: &ChanReader{
				PhaseSetting: sequence[i],
				InputSignal:  orchestrator.Senders[id],
			},
			Writer: &ChanWriter{
				OutputSignal: orchestrator.Receiver,
			},
			Program: &Program{
				Process: process,
			},
		}

		orchestrator.Amplifiers[i] = computer
	}
}

func getSignalTarget(signal Signal) string {
	var target string
	switch signal.Source {
	case "A":
		target = "B"
	case "B":
		target = "C"
	case "C":
		target = "D"
	case "D":
		target = "E"
	case "E":
		target = "A"
	}
	return target
}