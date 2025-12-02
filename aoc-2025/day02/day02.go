package main

import (
	"fmt"

	"github.com/trekawek/advent-of-code/aoc-2025/day02/part01"
	"github.com/trekawek/advent-of-code/aoc-2025/day02/part02"
	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func main() {
	lines := utils.ReadAllLines("input.txt")
	ranges := part01.ParseLines(lines)

	var result uint64
	switch utils.GetPartNumber() {
	case 1:
		result = part01.Solve(ranges)
	case 2:
		result = part02.Solve(ranges)
	default:
		panic(fmt.Errorf("invalid part number: %d", result))
	}

	fmt.Println(result)
}
