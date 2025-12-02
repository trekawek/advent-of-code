package part02

import (
	"testing"

	"github.com/trekawek/advent-of-code/aoc-2025/day02/part01"
	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func TestSolve(t *testing.T) {
	lines := utils.ReadAllLines("../test.txt")
	ranges := part01.ParseLines(lines)

	got := Solve(ranges)

	want := uint64(4174379265)
	if got != want {
		t.Errorf("want: %d got: %d", want, got)
	}

}

func TestIsInvalid_ForInvalidValue_ReturnsTrue(t *testing.T) {
	var input uint64
	input = 112233112233

	got := IsInvalid(input)
	want := true

	if got != want {
		t.Errorf("IsInvalid(%d) == %v, want: %v", input, got, want)
	}
}
