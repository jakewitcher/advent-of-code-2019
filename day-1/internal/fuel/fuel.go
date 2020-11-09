package fuel

func Calculate(mass int) int {
	return mass/3 - 2
}

func CalculateAll(mass int) int {
	var total int
	for {
		fuel := Calculate(mass)
		if fuel <= 0 {
			break
		}

		total += fuel
		mass = fuel
	}

	return total
}
