package intcode

func PowerThrusters(process []int, sequence []int) (int, error) {
	writer := MemWriter{}

	for _, phase := range sequence {
		p := make([]int, len(process))
		copy(p, process)

		reader := MemReader{}

		reader.PhaseSetting = phase
		reader.InputSignal = writer.OutputSignal

		intCodeComputer := IntCodeComputer{
			Reader: &reader,
			Writer: &writer,
			Program: &Program{Process: p},
		}

		err := intCodeComputer.Run()

		if err != nil {
			return 0, err
		}
	}

	return writer.OutputSignal, nil
}
