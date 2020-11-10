package wires

import (
	"day-3/internal/wires"
	"testing"
)

var testCases = []struct {
	wireA, wireB    []string
	distance, steps int
}{
	{
		[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
		159,
		610,
	},
	{
		[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
		135,
		410,
	},
}

func TestClosestIntersection(t *testing.T) {
	for _, test := range testCases {
		if distance := wires.ClosestIntersection(test.wireA, test.wireB); distance != test.distance {
			t.Fatalf("expected %v\ngot %v", test.distance, distance)
		}
	}
}

func TestShortestDistance(t *testing.T) {
	for _, test := range testCases {
		if steps := wires.ShortestSteps(test.wireA, test.wireB); steps != test.steps {
			t.Fatalf("expected %v\ngot %v", test.steps, steps)
		}
	}
}
