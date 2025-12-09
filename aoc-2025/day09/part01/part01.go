package part01

import (
	"fmt"
	"slices"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

type Pair struct {
	x int
	y int
}

func Area(p1 Pair, p2 Pair) int {
	return utils.Abs(p1.x-p2.x+1) * utils.Abs(p1.y-p2.y+1)
}

func Parse(lines []string) []Pair {
	var pairs []Pair
	for _, l := range lines {
		p := Pair{}
		fmt.Sscanf(l, "%d,%d", &p.x, &p.y)
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
