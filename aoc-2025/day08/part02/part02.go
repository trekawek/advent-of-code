package part02

import (
	"github.com/trekawek/advent-of-code/aoc-2025/day08/part01"
)

func Solve(lines []string) int {
	boxes := part01.Parse(lines)
	sets := part01.InitializeSets(len(boxes))
	pairs := part01.GetSortedPairs(boxes)

	var result int
	for _, p := range pairs {
		if part01.Merge(sets, p.B1, p.B2) {
			result = boxes[p.B1][0] * boxes[p.B2][0]
		}
	}
	return result
}
