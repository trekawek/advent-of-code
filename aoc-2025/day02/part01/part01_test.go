package part01

import (
	"testing"

	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func TestSolve(t *testing.T) {
	lines := utils.ReadAllLines("../test.txt")
	ranges := ParseLines(lines)

	got := Solve(ranges)

	want := uint64(1227775554)
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

func TestIsInvalid_ForValidValue_ReturnsFalse(t *testing.T) {
	var input uint64
	input = 1213

	got := IsInvalid(input)
	want := false

	if got != want {
		t.Errorf("IsInvalid(%d) == %v, want: %v", input, got, want)
	}
}

func TestIsInvalid_ForNonEvenLenghts_ReturnsFalse(t *testing.T) {
	var input uint64
	input = 121

	got := IsInvalid(input)
	want := false

	if got != want {
		t.Errorf("IsInvalid(%d) == %v, want: %v", input, got, want)
	}
}
