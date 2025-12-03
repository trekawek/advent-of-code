package part02

import (
	"testing"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func TestSolve(t *testing.T) {
	lines := utils.ReadAllLines("../test.txt")

	got := Solve(lines)

	want := uint64(3121910778619)
	if got != want {
		t.Errorf("want: %d got: %d", want, got)
	}

}
