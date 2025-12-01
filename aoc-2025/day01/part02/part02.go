package part02

import (
	"fmt"

	"github.com/trekawek/advent-of-code/aoc-2025/day01/part01"
)

func Solve(lines []string) int {
	s := part01.START
	result := 0

	for _, m := range part01.ParseMoves(lines) {
		var r int
		s, r = m.Apply(s)
		result += r

		fmt.Printf("%s\t%d\t%d\n", m.String(), s, r)
	}

	return result
}
