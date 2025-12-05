package main

import (
	"fmt"

	"github.com/trekawek/advent-of-code/aoc-2025/day05/part01"
	"github.com/trekawek/advent-of-code/aoc-2025/day05/part02"
	"github.com/trekawek/advent-of-code/aoc-2025/utils"
)

func main() {
	lines := utils.ReadAllLines("input.txt")

	var result int
	switch utils.GetPartNumber() {
	case 1:
		result = part01.Solve(lines)
	case 2:
		result = part02.Solve(lines)
	default:
		panic(fmt.Errorf("invalid part number: %d", result))
	}

	fmt.Println(result)
}
