package part02

import (
	"github.com/trekawek/advent-of-code/aoc-2025/day04/part01"
	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func Solve(lines []string) int {
	f := utils.ParseField(lines)
	c := 0

	for {
		updated := false
		for x, y := range f.Iterate {
			if f.Get(x, y) != '@' {
				continue
			}
			if part01.CountAdjacent(f, x, y) < 4 {
				updated = true
				c++
				f.Set(x, y, '.')
			}
		}
		if !updated {
			break
		}
	}
	return c
}
