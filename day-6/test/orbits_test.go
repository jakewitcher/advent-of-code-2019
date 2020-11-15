package orbits

import (
	"day-6/internal/orbits"
	"testing"
)

var totalOrbitsTestCases = []struct {
	orbitsMap []string
	total     int
}{
	{
		orbitsMap: []string{"COM)B", "B)C"},
		total:     3,
	},
	{
		orbitsMap: []string{"COM)B", "B)C", "B)D"},
		total:     5,
	},
	{
		orbitsMap: []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
		total:     42,
	},
}

var orbitalTransfersTestCases = []struct {
		orbitsMap []string
		total int
}{
	{
		orbitsMap: []string{"COM)B", "B)C", "B)YOU", "C)SAN"},
		total:     1,
	},
	{
		orbitsMap: []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
		total:     4,
	},
}

func TestTotalNumberOfOrbits(t *testing.T) {
	for _, test := range totalOrbitsTestCases {
		if total := orbits.TotalNumberOfOrbits(test.orbitsMap); total != test.total {
			t.Fatalf("expected %v, got %v", test.total, total)
		}
	}
}

func TestOrbitalTransfer(t *testing.T) {
	for _, test := range orbitalTransfersTestCases {
		if total, err := orbits.OrbitalTransfers(test.orbitsMap); total != test.total || err != nil {
			t.Fatalf("expected %v, got %v", test.total, total)
		}
	}
}
