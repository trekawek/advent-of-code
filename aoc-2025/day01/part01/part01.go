package part01

import (
	"fmt"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func (m *Move) Apply(v int) (int, int) {
	var w, r int
	if m.d == 'L' {
		w = v - m.v
		if w <= 0 {
			if v > 0 {
				r += 1
			}
			r += utils.Abs(w) / MAX
		}
	}
	if m.d == 'R' {
		w = v + m.v
		r += w / MAX
	}
	w = utils.Mod(w, MAX)
	return w, r
}

func Solve(lines []string) int {
	s := START
	result := 0

	for _, m := range ParseMoves(lines) {
		s, _ = m.Apply(s)
		if s == 0 {
			result += 1
		}

		fmt.Printf("%s\t%d\n", m.String(), s)
	}

	return result
}
