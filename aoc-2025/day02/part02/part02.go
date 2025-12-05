package part02

import (
	"fmt"
	"strings"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func IsInvalid(i uint64) bool {
	s := fmt.Sprintf("%d", i)
	l := len(s)

	k := 1
	for k <= l/2 {
		if l%k == 0 {
			t := strings.Repeat(s[:k], l/k)
			if t == s {
				return true
			}
		}
		k++
	}
	return false
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
