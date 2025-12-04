package part01

import "github.com/trekawek/advent-of-code/aoc-2025/utils"

func CountAdjacent(f *utils.Field, x int, y int) int {
	c := 0
	for i, j := range f.IterateAdjacent(x, y) {
		if f.Get(i, j) == '@' {
			c++
		}
	}
	return c
}

func Solve(lines []string) int {
	f := utils.ParseField(lines)
	c := 0
	for x, y := range f.Iterate {
		if f.Get(x, y) != '@' {
			continue
		}
		if CountAdjacent(f, x, y) < 4 {
			c++
		}
	}
	return c
}
