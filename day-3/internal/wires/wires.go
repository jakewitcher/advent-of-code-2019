package wires

import (
	"log"
	"math"
	"strconv"
)

type Point struct{ x, y float64 }

func ClosestIntersection(wireA, wireB []string) int {
	distance := math.MaxFloat64

	pointsA := getPoints(wireA)
	pointsB := getPoints(wireB)

	for point := range pointsA {
		if _, ok := pointsB[point]; ok {
			d := math.Abs(point.x) + math.Abs(point.y)
			if d < distance {
				distance = d
			}
		}
	}

	return int(distance)
}

func ShortestSteps(wireA, wireB []string) int {
	steps := math.MaxInt64

	pointsA := getPoints(wireA)
	pointsB := getPoints(wireB)

	for point, a := range pointsA {
		if b, ok := pointsB[point]; ok {
			s := a + b
			if s < steps {
				steps = s
			}
		}
	}

	return steps
}

func getPoints(wire []string) map[Point]int {
	points := make(map[Point]int)
	prev := Point{0, 0}
	count := 1

	for _, step := range wire {
		direction, length := parse(step)
		var move func(p Point) Point

		switch direction {
		case 'U':
			move = func(p Point) Point { return Point{p.x, p.y + 1} }
		case 'D':
			move = func(p Point) Point { return Point{p.x, p.y - 1} }
		case 'L':
			move = func(p Point) Point { return Point{p.x - 1, p.y} }
		case 'R':
			move = func(p Point) Point { return Point{p.x + 1, p.y} }
		default:
			log.Fatalf("invalid direction: %v", direction)
		}

		for i := 0; i < length; i++ {
			next := move(prev)
			if _, ok := points[next]; !ok {
				points[next] = count
			}

			prev = next
			count++
		}
	}

	return points
}

func parse(step string) (rune, int) {
	direction := rune(step[0])

	length, err := strconv.Atoi(step[1:])
	if err != nil {
		log.Fatalf("invalid length: %v", step[1:])
	}

	return direction, length
}
