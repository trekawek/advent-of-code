package part01

import (
	"fmt"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func IsInvalid(i uint64) bool {
	s := fmt.Sprintf("%d", i)
	if len(s)%2 == 1 {
		return false
	}
	m := len(s) / 2
	return s[:m] == s[m:]
}

func Solve(ranges []utils.Range) uint64 {
	var s uint64
	for _, r := range ranges {
		for i := range r.Iterate {
			if IsInvalid(i) {
				s += i
			}
		}
	}
	return s
}
