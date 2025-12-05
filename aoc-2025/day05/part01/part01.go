package part01

import (
	"strconv"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func Parse(lines []string) ([]utils.Range, []uint64) {
	var ranges []utils.Range
	var ids []uint64

	parsingRanges := true
	for _, l := range lines {
		if l == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			ranges = append(ranges, *utils.ParseRange(l))
		} else {
			i, err := strconv.ParseUint(l, 10, 64)
			if err != nil {
				panic(err)
			}
			ids = append(ids, i)
		}
	}
	return ranges, ids
}

func Solve(lines []string) int {
	c := 0
	ranges, ids := Parse(lines)
	for _, i := range ids {
		for _, r := range ranges {
			if r.Contains(i) {
				c++
				break
			}
		}
	}
	return c
}
