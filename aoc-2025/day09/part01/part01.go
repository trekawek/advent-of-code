package part01

import (
	"fmt"
	"slices"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

type Pair struct {
	X int
	Y int
}

func Area(p1 Pair, p2 Pair) int {
	return (utils.Abs(p1.X-p2.X) + 1) * (utils.Abs(p1.Y-p2.Y) + 1)
}

func Parse(lines []string) []Pair {
	var pairs []Pair
	for _, l := range lines {
		p := Pair{}
		fmt.Sscanf(l, "%d,%d", &p.X, &p.Y)
		pairs = append(pairs, p)
	}
	return pairs
}

func Solve(lines []string) int {
	pairs := Parse(lines)
	var areas []int
	for _, p1 := range pairs {
		for _, p2 := range pairs {
			a := Area(p1, p2)
			areas = append(areas, a)
		}
	}
	return slices.Max(areas)
}
