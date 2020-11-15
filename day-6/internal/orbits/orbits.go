package orbits

import (
	"errors"
	"strings"
)

type Body struct {
	self string
	orbits []*Body
}

func TotalNumberOfOrbits(input []string) int {
	bodies := getOrbits(input)
	return countOrbits(bodies, "COM", 1)
}

func OrbitalTransfers(input []string) (int, error) {
	bodies := getOrbits(input)
	you := getOrbitalPath(bodies, "COM", "YOU")
	san := getOrbitalPath(bodies, "COM", "SAN")

	return countTransfers(you, san)
}

func countTransfers(you []string, san []string) (int, error) {
	for i, node := range you {
		if node == san[i] && node == "YOU" {
			return len(san[i:]) - 1, nil
		}

		if node == san[i] && node == "SAN" {
			return len(you[i:]) - 1, nil
		}

		if node != san[i] {
			return len(you[i:]) + len(san[i:]) - 2, nil
		}
	}

	return 0, errors.New("no path available between YOU and SAN")
}

func getOrbitalPath(bodies map[string][]string, center string, target string) []string {
	path := []string{center}
	for _, orbit := range bodies[center] {
		if orbit == target {
			return append(path, orbit)
		}

		tail := getOrbitalPath(bodies, orbit, target)
		if tail == nil {
			continue
		}

		return append(path, tail...)
	}

	return nil
}

func getOrbits(input []string) map[string][]string {
	bodies := make(map[string][]string)

	for _, val := range input {
		body, orbit := parseBody(val)
		if _, ok := bodies[body]; ok {
			bodies[body] = append(bodies[body], orbit)
		} else {
			bodies[body] = []string{orbit}
		}
	}

	return bodies
}

func countOrbits(bodies map[string][]string, center string, depth int) int {
	var total int

	if orbits, ok := bodies[center]; ok {
		total += len(orbits) * depth

		for _, body := range orbits {
			total += countOrbits(bodies, body, depth + 1)
		}
	}

	return total
}

func parseBody(raw string) (string, string) {
	v := strings.Split(raw, ")")
	return v[0], v[1]
}
