package part02

import (
	"github.com/trekawek/advent-of-code/aoc-2025/day05/part01"
	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func overlaps(r, q utils.Range) bool {
	return r.Contains(q.From) || r.Contains(q.To) || q.Contains(r.From) || q.Contains(r.To)
}

func merge(r, q utils.Range) utils.Range {
	if r.Contains(q.From) && r.Contains(q.To) {
		return r
	} else if r.Contains(q.From) {
		return utils.NewRange(r.From, q.To)
	} else if r.Contains(q.To) {
		return utils.NewRange(q.From, r.To)
	}
	return merge(q, r)
}

func Solve(lines []string) uint64 {
	ranges, _ := part01.Parse(lines)

	for {
		updated := false
		var newRanges []utils.Range
		merged := map[int]bool{}
		for i, r := range ranges {
			if merged[i] {
				continue
			}
			for j, q := range ranges {
				if i == j {
					continue
				}
				if overlaps(r, q) {
					r = merge(r, q)
					updated = true
					merged[j] = true
				}
			}
			newRanges = append(newRanges, r)
		}

		ranges = newRanges
		if !updated {
			break
		}
	}

	var c uint64
	for _, r := range ranges {
		c += r.Length()
	}
	return c
}
